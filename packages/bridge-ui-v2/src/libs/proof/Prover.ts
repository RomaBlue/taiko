import { getContract, type GetContractResult, type PublicClient } from '@wagmi/core';
import { type Address, encodePacked, type Hex, keccak256, toHex } from 'viem';

import { crossChainSyncABI } from '$abi';
import { routingContractsMap } from '$bridgeConfig';
import { PendingBlockError } from '$libs/error';
import { getLogger } from '$libs/util/logger';
import { publicClient } from '$libs/wagmi';

import { type Block, type ClientWithEthGetProofRequest, type GenerateProofArgs, ProofAction } from './types';

const log = getLogger('proof:Prover');

export class Prover {
  protected async _getKeyToClaim(contractAddress: Address, msgHash: Hex) {
    return keccak256(encodePacked(['address', 'bytes32'], [contractAddress, msgHash]));
  }

  protected async _getKeyToRelease(msgHash: Hex) {
    return keccak256(encodePacked(['bytes', 'bytes32'], [toHex('MESSAGE_STATUS'), msgHash]));
  }

  protected async _getLatestBlock(
    client: PublicClient,
    crossChainSyncContract: GetContractResult<typeof crossChainSyncABI>,
  ) {
    const latestBlockHash = await crossChainSyncContract.read.getCrossChainBlockHash([BigInt(0)]);

    const block: Block = await client.request({
      method: 'eth_getBlockByHash',
      params: [
        latestBlockHash, false
      ],
    });
    return block;
  }

  async generateProof(args: GenerateProofArgs) {
    const { action, msgHash, srcChainId, contractAddress, destChainId, proofForAccountAddress } = args;

    let key: Hex = toHex(0);
    let client;

    if (action === ProofAction.CLAIM) {
      client = publicClient({ chainId: srcChainId });
      key = await this._getKeyToClaim(contractAddress, msgHash);
    } else if (action === ProofAction.RELEASE) {
      client = publicClient({ chainId: destChainId });
      key = await this._getKeyToRelease(msgHash);
    }

    if (!client) {
      throw new Error('client is not defined');
    }

    // Unfortunately, since this method is stagnant, it hasn't been included into Viem lib
    // as supported methods. Still stupported  by Alchmey, Infura and others.
    // See https://eips.ethereum.org/EIPS/eip-1186
    // Following is a workaround to support this method.
    const clientWithEthProofRequest = client as ClientWithEthGetProofRequest;

    const crossChainSyncAddress = routingContractsMap[destChainId][srcChainId].crossChainSyncAddress;

    // Get the block from chain A based on the latest block hash
    // we get cross chain (Taiko contract on chain B)
    const crossChainSyncContract = getContract({
      chainId: destChainId,
      address: crossChainSyncAddress,
      abi: crossChainSyncABI,
    });

    const client2 = publicClient({ chainId: srcChainId });
    const block = await this._getLatestBlock(client2, crossChainSyncContract);

    if (block.hash === null || block.number === null) {
      throw new PendingBlockError('block is pending');
    }
    // const block = null;
    // RPC call to get the merkle proof what value is at key on the SignalService contract
    const proof = await clientWithEthProofRequest.request({
      method: 'eth_getProof',
      params: [
        // Address of the account to get the proof for
        proofForAccountAddress,

        // Array of storage-keys that should be proofed and included
        [key],

        block.hash
      ],
    });

    log('Proof from eth_getProof', proof);

    return { proof, block };
  }
}
