// SPDX-License-Identifier: MIT

// 1. Specify the compiler version
pragma solidity ^0.8.20;

// 2. Define the contract (similar to a 'class' in other languages)
contract SimpleStorage {
    
    // 3. State Variable: This is permanently stored on the blockchain!
    // uint256 means "unsigned integer of 256 bits"
    uint256 public storedData;

    // 4. Write Function: Modifies the blockchain state. 
    // Executing this requires a transaction and costs "Gas" (transaction fee).
    function set(uint256 x) public {
        storedData = x;
    }

    // 5. Read Function: Does not modify state.
    // The 'view' keyword means it just reads data. It is completely free to execute.
    function get() public view returns (uint256) {
        return storedData;
    }
}
