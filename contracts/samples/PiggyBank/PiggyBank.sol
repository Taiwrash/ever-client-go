pragma solidity >=0.5.0;

contract PiggyBank {

    // State variables:
    address owner;        // contract owner's address;
    uint limit;            // piggybank's minimal limit to withdraw;
    uint128 balance;            // piggybank's deposit balance;
    uint version = 1;        // version of the PiggyBank.

    // Modifier that allows public function to accept all external calls.
    modifier alwaysAccept {
        // Runtime function that allows contract to process inbound messages spending
        // its own resources (it's necessary if contract should process all inbound messages,
        // not only those that carry value with them).
        tvm.accept();
        _;
    }

    // Modifier that allows public function to be called only by message signed with owner's pubkey.
    modifier checkPubkeyAndAccept {
        require(msg.pubkey() == tvm.pubkey());
        tvm.accept();
        _;
    }

    // Modifier that allows public function to be called only from the owners address.
    modifier checkOwnerAndAccept {
        require(msg.sender == owner);
        tvm.accept();
        _;
    }

    // Modifier that allows public function to be called only when the limit is reached.
    modifier checkBalance() {
        require(balance >= limit);
        _;
    }

    // Constructor saves the address of the contract owner in a state variable and
    // initializes the limit and the balance.
    constructor(address pb_owner, uint pb_limit) public alwaysAccept {
        owner = pb_owner;
        limit = pb_limit;
        balance = 0;
    }

    // Function to obtain parameters of the contract.
	function getData() public alwaysAccept returns (address, uint, uint128) {
		return (owner, limit, balance);
	}

    // Function that can be called by anyone.
    function deposit(uint amount) public alwaysAccept {
        balance += uint128(amount);
    }

    // Function that can be called only by the owner after reaching the limit.
    function withdraw() public checkBalance checkOwnerAndAccept {
        msg.sender.transfer(balance);
        balance = 0;
    }

    // Function to change the PiggyBank limit that can be called only by the owner
    function setLimit(uint newLimit) public checkOwnerAndAccept {
        limit = newLimit;
    }

    // Function that changes the code of current contract.
    function setCode(TvmCell newCode) public checkPubkeyAndAccept {
        // Runtime function that creates an output action that would change this
        // smart contract code to that given by cell newcode.
        tvm.setcode(newCode);

        // Runtime function that replaces current code of the contract with newcode.
        tvm.setCurrentCode(newCode);

        // Call function onCodeUpgrade of the 'new' code.
        onCodeUpgrade();
    }

    // After code upgrade caused by calling setCode function we may need to do some actions.
    // We can add them into this function with constant id.
    function onCodeUpgrade() private {
        version = 2;
        setLimit(limit * 2);
    }

    // Function to obtain contract version.
    function getVersion() public alwaysAccept returns (uint) {
        return version;
    }

}
