// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MtkContracts {
    //枚举
    enum StakingPeriod { ThirtyDays, NinetyDays, HundredEightyDays, OneYear }
    mapping(StakingPeriod => uint256) public apy; // 年化收益率（百分比，如20表示20%）
    
    //定义结构体
    struct Stake{
        uint256 stakeIndex;         //唯一索引
        uint256 amount;             // 质押数量
        uint256 startTime;          // 质押开始时间
        uint256 endTime;            // 质押结束时间
        uint256 rewardRate;         // 收益率（根据期限计算）
        bool isActive;              // 订单是否有效
    }
    

    IERC20 public  stakingToken;//
 
    mapping(address => Stake[]) public userStakes; // 用户的所有质押订单
    mapping(StakingPeriod => uint256) public aprs; // 不同期限（秒）对应的年化收益率
    mapping(StakingPeriod=>uint256) public durations; // 不同期限对应的秒数 )

     event Staked(
        address indexed user,
        uint256 amount,
        StakingPeriod period,
        uint256 timestamp
    );

    event Withdrawn(
        address indexed user,
        uint256 totalAmount,
        uint256 stakeIndex
    );

    constructor(IERC20 _mtkToken) {
        stakingToken = _mtkToken;

        durations[StakingPeriod.ThirtyDays] = 30 days;
        durations[StakingPeriod.NinetyDays] = 90 days;
        durations[StakingPeriod.HundredEightyDays] = 180 days;
        durations[StakingPeriod.OneYear] = 365 days;

        apy[StakingPeriod.ThirtyDays] = 10;   // 10% 年化
        apy[StakingPeriod.NinetyDays] = 15;   // 15% 年化
        apy[StakingPeriod.HundredEightyDays] = 18; // 18% 年化
        apy[StakingPeriod.OneYear] = 20;      // 20% 年化
    }

    function stake(uint256 amount, StakingPeriod period) external {
       
        require(amount>0,"amount must be greater than zero");
        require(stakingToken.transferFrom(msg.sender, address(this), amount), "Transfer failed");

        uint256 duration=_getDuration(period);//获取分钟
        uint256 periodDays=durations[period];
        uint256 rate=apy[period]*periodDays*1**18/360;
        uint256 start = block.timestamp;//区块开始时间
        uint256 end = start + duration;//结束时间
        
        Stake memory newStake = Stake({
            stakeIndex: userStakes[msg.sender].length,
            amount: amount,
            startTime: start,
            endTime: end,
            rewardRate: rate, // 收益率（根据期限计算）
            isActive: true    // 订单是否有效
        });

        userStakes[msg.sender].push(newStake);
        //触发事件
        emit Staked(msg.sender, amount, period, end);
    }

    // 内部函数：根据期限返回秒数
    function _getDuration(StakingPeriod period) internal pure returns (uint256) {
        if(period==StakingPeriod.ThirtyDays){
            return  1 minutes;
        }else if(period==StakingPeriod.NinetyDays){
            return  3 minutes;
        }else if(period==StakingPeriod.HundredEightyDays){
            return  5 minutes;
        }else {
            return  10 minutes;
        }
    }
        
    //提现
    function withdraw(uint256 stakeIndex) external {
        require(stakeIndex<userStakes[msg.sender].length,"Invalid stake index");
        Stake storage stk = userStakes[msg.sender][stakeIndex];
        require(stk.isActive, "stk is not active");
        //质押时间没有到
        require(block.timestamp >= stk.endTime, "Staking period is not over");
        stk.isActive=false;
        
        uint256 reward = stk.amount * stk.rewardRate  / 100/10**18;
        uint256 totalAmount = stk.amount + reward;
        // 将质押的代币和收益转移给用户
        stakingToken.transfer(msg.sender, totalAmount);
        emit Withdrawn(msg.sender, totalAmount, stakeIndex);
    }
   


}
