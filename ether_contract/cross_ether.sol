pragma solidity ^0.6.0;
// SPDX-License-Identifier: UNLICENSED

/**
 * @title ERC20 interface
 * @dev see https://eips.ethereum.org/EIPS/eip-20
 */
interface InterfaceERC20 {
    function totalSupply() external view returns (uint256);
    function transfer(address to, uint256 value) external returns (bool);
    function approve(address spender, uint256 value) external returns (bool);
    function transferFrom(address from, address to, uint256 value) external returns (bool);
    function balanceOf(address who) external view returns (uint256);
    function allowance(address owner, address spender) external view returns (uint256);
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}

/**
 * @title SafeMath
 * @dev Unsigned math operations with safety checks that revert on error.
 */
library SafeMath {
    /**
     * @dev Multiplies two unsigned integers, reverts on overflow.
     */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
        if (a == 0) {
            return 0;
        }
        uint256 c = a * b;
        require(c / a == b);
        return c;
    }

    /**
     * @dev Integer division of two unsigned integers truncating the quotient, reverts on division by zero.
     */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        // Solidity only automatically asserts when dividing by 0
        require(b > 0);
        uint256 c = a / b;
        // assert(a == b * c + a % b); // There is no case in which this doesn't hold
        return c;
    }

    /**
     * @dev Subtracts two unsigned integers, reverts on overflow (i.e. if subtrahend is greater than minuend).
     */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a);
        uint256 c = a - b;
        return c;
    }

    /**
     * @dev Adds two unsigned integers, reverts on overflow.
     */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a);
        return c;
    }

    /**
     * @dev Divides two unsigned integers and returns the remainder (unsigned integer modulo),
     * reverts when dividing by zero.
     */
    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b != 0);
        return a % b;
    }
}

contract CrossEther {
    using SafeMath for uint256;
    mapping(address => bool) owner;
    uint256 public total;
    // ETH => TRX
    mapping(address => string) ethUsers;
    // TRX => ETH
    mapping(string => address) trxUsers;
    // 用户当前余额
    mapping(address => uint256) userTotal;

    address public ERC20Address;
    event Receive(address addr,uint256 value);
    event Withdraw(address addr,uint256 value);

    constructor(address erc20) public{
        owner[msg.sender] = true;
        ERC20Address = erc20;
    }

    // 充值
    function recharge(uint256 value) public{
        require(value > 0);
        InterfaceERC20(ERC20Address).transferFrom(msg.sender,address(this),value);
        total = total.add(value);
        userTotal[msg.sender] = userTotal[msg.sender].add(value);
        emit Receive(msg.sender,value);
    }

    // 管理员提币
    function ownerWithdraw(address payable addr,uint256 value) public onlyOwner{
        require(value != 0 && addr != address(0));
        InterfaceERC20(ERC20Address).transfer(addr,value);
        emit Withdraw(addr,value);
    }

    // 地址绑定
    function setTrxAddress(string memory addr) public {
        require(trxUsers[addr] == address(0));
        ethUsers[msg.sender] = addr;
        trxUsers[addr] = msg.sender;
    }

    function getUserTrxByEthAddr(address addr) view public returns(string memory){
        return ethUsers[addr];
    }

    function getUserEthByTrxAddr(string memory addr) view public returns(address){
        return trxUsers[addr];
    }

    // 管理员权限转移
    function ownerPermissions(address newOwner,bool isOwner) public onlyOwner{
        require(newOwner != address(0));
        owner[newOwner] = isOwner;
    }

    // 更改ERC20地址
    function ownerSetErc20Addr(address addr) public onlyOwner{
        require(addr != address(0));
        ERC20Address = addr;
    }

    modifier onlyOwner() {
        require(owner[msg.sender] == true);
        _;
    }
}
