import { encodeAbiParameters, encodePacked, type Hash, type Hex, toHex, toRlp } from 'viem';

import { routingContractsMap } from '$bridgeConfig';
import { MessageStatus } from '$libs/bridge';
import { InvalidProofError } from '$libs/error';

import { Prover } from './Prover';
import { type Block, type BlockHeader, type EthGetProofResponse, ProofAction } from './types';

export class BridgeProver extends Prover {
  constructor() {
    super();
  }

  private _getSignalProof(proof: EthGetProofResponse, blockHeight: bigint) {
    // RLP encode the proof together for LibTrieProof to decode
    const encodedProof = toRlp(proof.storageProof[0].proof);

    // Encode the SignalProof struct:
    // struct SignalProof {
    //   uint256 height;
    //   bytes proof;
    // }
    const signalProof = encodeAbiParameters(
      // ['tuple(uint256 height, bytes proof)'],
      [
        {
          type: 'tuple',
          components: [
            { name: 'height', type: 'uint256' },
            { name: 'proof', type: 'bytes' },
          ],
        },
      ],
      [{ height: blockHeight, proof: encodedProof }],
    );

    return signalProof;
  }

  private _getSignalProofForRelease(proof: EthGetProofResponse, blockHeader: BlockHeader) {
    // RLP encode the proof together for LibTrieProof to decode
    const encodedAccountProof = toRlp(proof.accountProof[0]);
    const encodedStorageProof = toRlp(proof.storageProof[0].proof);

    // RLP encode the proof together for LibTrieProof to decode
    // const encodedProof = ethers.utils.defaultAbiCoder.encode(
    //   ['bytes', 'bytes'],
    //   [RLP.encode(proof.accountProof), RLP.encode(proof.storageProof[0].proof)],
    // );

    // encode the SignalProof struct from LibBridgeSignal
    // const signalProof = ethers.utils.defaultAbiCoder.encode(
    //   [
    //     'tuple(tuple(bytes32 parentHash, bytes32 ommersHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes32[8] logsBloom, uint256 difficulty, uint128 height, uint64 gasLimit, uint64 gasUsed, uint64 timestamp, bytes extraData, bytes32 mixHash, uint64 nonce, uint256 baseFeePerGas, bytes32 withdrawalsRoot) header, bytes proof)',
    //   ],
    //   [{ header: blockHeader, proof: encodedProof }],
    // );

    const encodedProof = encodePacked(
      ['bytes', 'bytes'],
      [encodedAccountProof, encodedStorageProof],
    );

    const signalProof = encodePacked(
      [
        'tuple(tuple(bytes32 parentHash, bytes32 ommersHash, address beneficiary, bytes32 stateRoot, bytes32 transactionsRoot, bytes32 receiptsRoot, bytes32[8] logsBloom, uint256 difficulty, uint128 height, uint64 gasLimit, uint64 gasUsed, uint64 timestamp, bytes extraData, bytes32 mixHash, uint64 nonce, uint256 baseFeePerGas, bytes32 withdrawalsRoot) header, bytes proof)',
      ],
      [{ header: blockHeader, proof: encodedProof }],
    );

    return signalProof;
  }

  async generateProofToProcessMessage(msgHash: Hash, srcChainId: number, destChainId: number) {
    const srcBridgeAddress = routingContractsMap[srcChainId][destChainId].bridgeAddress;
    const srcSignalServiceAddress = routingContractsMap[srcChainId][destChainId].signalServiceAddress;

    const { proof, block } = await this.generateProof({
      action: ProofAction.CLAIM,
      msgHash,
      srcChainId,
      contractAddress: srcBridgeAddress,
      destChainId,
      proofForAccountAddress: srcSignalServiceAddress,
    });

    // Value must be 0x1 => isSignalSent
    if (proof.storageProof[0].value !== toHex(true)) {
      throw new InvalidProofError('storage proof value is not 1');
    }

    return this._getSignalProof(proof, BigInt(block.number));
  }



  async generateProofToRelease(msgHash: Hash, srcChainId: number, destChainId: number) {
    // const srcBridgeAddress = routingContractsMap[srcChainId][destChainId].bridgeAddress;
    const destBridgeAddress = routingContractsMap[destChainId][srcChainId].bridgeAddress;

    const { proof, block } = await this.generateProof({
      action: ProofAction.RELEASE,
      msgHash,
      srcChainId,
      contractAddress: destBridgeAddress,
      destChainId,
      proofForAccountAddress: destBridgeAddress,
    });

    const blockHeader = buildBlockHeaderFromBlock(block);

    // eslint-disable-next-line no-console
    console.log('Proof from eth_getProof', proof);
    // Value must be 0x3 => MessageStatus.FAILED
    if (proof.storageProof[0].value !== toHex(MessageStatus.FAILED)) {
      throw new InvalidProofError('storage proof value is not FAILED');
    }

    return this._getSignalProofForRelease(proof, blockHeader)
  }
}
const buildBlockHeaderFromBlock = (block: Block): BlockHeader => {
  const {
    parentHash,
    sha3Uncles,
    miner,
    stateRoot,
    transactionsRoot,
    receiptsRoot,
    logsBloom,
    difficulty,
    number,
    gasLimit,
    gasUsed,
    timestamp,
    extraData,
    mixHash,
    baseFeePerGas,
    withdrawalsRoot
  } = block;

  // const logsBloominput: `0x${string}` | `0x${string}`[] = logsBloom;

  // const logsBloomString = logsBloominput.toString().substring(2);
  // const logsBloomArray = logsBloomString.match(/.{1,64}/g)!.map((s: string) => '0x' + s) as Hex[];

  let logsBloomArray: Hex[];

  if (Array.isArray(logsBloom)) {
    logsBloomArray = logsBloom;
  } else {
    const logsBloomString = logsBloom.substring(2);
    logsBloomArray = logsBloomString.match(/.{1,64}/g)!.map((s) => `0x${s}`) as Hex[];
  }

  return {
    parentHash,
    ommersHash: sha3Uncles,
    proposer: miner,
    stateRoot,
    transactionsRoot,
    receiptsRoot,
    logsBloom: logsBloomArray,
    difficulty,
    height: number ? toHex(number) : toHex(0),
    gasLimit,
    gasUsed,
    timestamp,
    extraData,
    mixHash,
    nonce: toHex(0),
    baseFeePerGas,
    withdrawalsRoot
  };
};


//   struct BlockHeader {
//     bytes32 transactionsRoot;
//     bytes32 receiptsRoot;
//     bytes32[8] logsBloom;
//     uint256 difficulty;
//     uint128 height;
//     uint64 gasLimit;
//     uint64 gasUsed;
//     uint64 timestamp;
//     bytes extraData;
//     bytes32 mixHash;
//     uint64 nonce;
//     uint256 baseFeePerGas;
//     bytes32 withdrawalsRoot;
// }