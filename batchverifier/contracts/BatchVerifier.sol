
pragma solidity ^0.8.13;

contract BatchVerifier {


	mapping (address => uint) balances;

	event Transfer(address indexed _from, address indexed _to, uint256 _value);

}
