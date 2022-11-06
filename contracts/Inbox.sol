// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.4.21 <0.8.18;

contract Inbox {
    string public message;

    constructor(string memory initialMessage) public {
        message = initialMessage;
    }

    function setMessage(string memory newMessage) public {
        message = newMessage;
    }
}
