import type { Hex } from "viem";

export const generateZeroHex = (length: number): Hex => {
    return `0x${'0'.repeat(length * 2)}`;
};