// Generated by github.com/davyxu/tabtoy
// Version: 2.8.10

module table {
export var ProtoId : table.IProtoIdDefine[] = [
		{ Id : 1, Name : "msg.AccountInfo" 	},
		{ Id : 2, Name : "msg.AccountGateInfo" 	},
		{ Id : 3, Name : "msg.GridItem" 	},
		{ Id : 4, Name : "msg.BT_UploadGameUser" 	},
		{ Id : 5, Name : "msg.BT_ReqEnterRoom" 	},
		{ Id : 6, Name : "msg.BT_GameInit" 	},
		{ Id : 7, Name : "msg.BT_GameStart" 	},
		{ Id : 8, Name : "msg.BT_GameEnd" 	},
		{ Id : 9, Name : "msg.BT_GameOver" 	},
		{ Id : 10, Name : "msg.BT_ReqQuitGameRoom" 	},
		{ Id : 11, Name : "msg.BT_PickItem" 	},
		{ Id : 12, Name : "msg.BT_ReqLaunchBullet" 	},
		{ Id : 13, Name : "msg.BT_RetLaunchBullet" 	},
		{ Id : 14, Name : "msg.BT_StepOnBomb" 	},
		{ Id : 15, Name : "msg.BT_RetStepOnBomb" 	},
		{ Id : 16, Name : "msg.BT_BulletEarnMoney" 	},
		{ Id : 17, Name : "msg.BT_UseUltimateSkil" 	},
		{ Id : 18, Name : "msg.BT_ReqCrushSuperBrick" 	},
		{ Id : 19, Name : "msg.BT_RetCrushSuperBrick" 	},
		{ Id : 20, Name : "msg.C2GW_JoinGame" 	},
		{ Id : 21, Name : "msg.GW2C_JoinOk" 	},
		{ Id : 22, Name : "msg.RoomMemberInfo" 	},
		{ Id : 23, Name : "msg.GW2C_UpdateRoomInfo" 	},
		{ Id : 24, Name : "msg.GW2C_StartGame" 	},
		{ Id : 25, Name : "msg.GW2C_QuestionInfo" 	},
		{ Id : 26, Name : "msg.C2GW_Answer" 	},
		{ Id : 27, Name : "msg.GW2C_AnswerOk" 	},
		{ Id : 28, Name : "msg.GW2C_AnswerInfo" 	},
		{ Id : 29, Name : "msg.GW2C_GameOver" 	},
		{ Id : 30, Name : "msg.C2GW_ReqLogin" 	},
		{ Id : 31, Name : "msg.GW2C_RetLogin" 	},
		{ Id : 32, Name : "msg.GW2C_SendUserInfo" 	},
		{ Id : 33, Name : "msg.GW2C_SendUserPlatformMoney" 	},
		{ Id : 34, Name : "msg.C2GW_HeartBeat" 	},
		{ Id : 35, Name : "msg.GW2C_HeartBeat" 	},
		{ Id : 36, Name : "msg.C2GW_ReqStartGame" 	},
		{ Id : 37, Name : "msg.GW2C_RetStartGame" 	},
		{ Id : 38, Name : "msg.GW2C_Ret7DayReward" 	},
		{ Id : 39, Name : "msg.C2GW_Get7DayReward" 	},
		{ Id : 40, Name : "msg.C2GW_SendWechatAuthCode" 	},
		{ Id : 41, Name : "msg.C2GW_GoldExchange" 	},
		{ Id : 42, Name : "msg.GW2C_RetGoldExchange" 	},
		{ Id : 43, Name : "msg.C2L_ReqLogin" 	},
		{ Id : 44, Name : "msg.C2L_ReqLoginWechat" 	},
		{ Id : 45, Name : "msg.L2C_RetLogin" 	},
		{ Id : 46, Name : "msg.C2L_ReqRegistAuthCode" 	},
		{ Id : 47, Name : "msg.C2L_ReqRegistAccount" 	},
		{ Id : 48, Name : "msg.L2C_RetRegistAccount" 	},
		{ Id : 49, Name : "msg.IpHost" 	},
		{ Id : 50, Name : "msg.PairNumItem" 	},
		{ Id : 51, Name : "msg.GW2L_ReqRegist" 	},
		{ Id : 52, Name : "msg.L2GW_RetRegist" 	},
		{ Id : 53, Name : "msg.GW2L_HeartBeat" 	},
		{ Id : 54, Name : "msg.L2GW_HeartBeat" 	},
		{ Id : 55, Name : "msg.L2GW_ReqRegistUser" 	},
		{ Id : 56, Name : "msg.GW2L_RegistUserRet" 	},
		{ Id : 57, Name : "msg.GW2MS_ReqRegist" 	},
		{ Id : 58, Name : "msg.MS2GW_RetRegist" 	},
		{ Id : 59, Name : "msg.GW2MS_HeartBeat" 	},
		{ Id : 60, Name : "msg.MS2GW_HeartBeat" 	},
		{ Id : 61, Name : "msg.GW2MS_ReqCreateRoom" 	},
		{ Id : 62, Name : "msg.MS2GW_RetCreateRoom" 	},
		{ Id : 63, Name : "msg.RS2GW_ReqRegist" 	},
		{ Id : 64, Name : "msg.GW2RS_RetRegist" 	},
		{ Id : 65, Name : "msg.GW2RS_UserDisconnect" 	},
		{ Id : 66, Name : "msg.RS2GW_RetUserDisconnect" 	},
		{ Id : 67, Name : "msg.GW2RS_MsgTransfer" 	},
		{ Id : 68, Name : "msg.RS2GW_MsgTransfer" 	},
		{ Id : 69, Name : "msg.C2GW_BuyItem" 	},
		{ Id : 70, Name : "msg.GW2C_AddPackageItem" 	},
		{ Id : 71, Name : "msg.GW2C_RemovePackageItem" 	},
		{ Id : 72, Name : "msg.GW2C_UpdateGold" 	},
		{ Id : 73, Name : "msg.GW2C_UpdateYuanbao" 	},
		{ Id : 74, Name : "msg.GW2C_UpdateDiamond" 	},
		{ Id : 75, Name : "msg.GW2C_UpdateFreeStep" 	},
		{ Id : 76, Name : "msg.DeliveryGoods" 	},
		{ Id : 77, Name : "msg.C2GW_ReqDeliveryGoods" 	},
		{ Id : 78, Name : "msg.BigRewardItem" 	},
		{ Id : 79, Name : "msg.Sync_BigRewardPickNum" 	},
		{ Id : 80, Name : "msg.C2GW_UseBagItem" 	},
		{ Id : 81, Name : "msg.C2GW_SellBagItem" 	},
		{ Id : 82, Name : "msg.C2GW_BuyClothes" 	},
		{ Id : 83, Name : "msg.C2GW_DressClothes" 	},
		{ Id : 84, Name : "msg.C2GW_UnDressClothes" 	},
		{ Id : 85, Name : "msg.GW2C_UpdateItemPos" 	},
		{ Id : 86, Name : "msg.GW2C_SendShowImage" 	},
		{ Id : 87, Name : "msg.C2GW_ChangeImageSex" 	},
		{ Id : 88, Name : "msg.GW2C_RetChangeImageSex" 	},
		{ Id : 89, Name : "msg.RS2MS_ReqRegist" 	},
		{ Id : 90, Name : "msg.MS2RS_RetRegist" 	},
		{ Id : 91, Name : "msg.RS2MS_HeartBeat" 	},
		{ Id : 92, Name : "msg.MS2RS_HeartBeat" 	},
		{ Id : 93, Name : "msg.GateSimpleInfo" 	},
		{ Id : 94, Name : "msg.MS2RS_GateInfo" 	},
		{ Id : 95, Name : "msg.MS2RS_CreateRoom" 	},
		{ Id : 96, Name : "msg.RS2MS_RetCreateRoom" 	},
		{ Id : 97, Name : "msg.RS2MS_DeleteRoom" 	},
		{ Id : 98, Name : "msg.RS2MS_UpdateRewardPool" 	},
		{ Id : 99, Name : "msg.GW2C_MsgNotify" 	},
		{ Id : 100, Name : "msg.GW2C_MsgNotice" 	},
		{ Id : 101, Name : "msg.GW2MS_MsgNotice" 	},
		{ Id : 102, Name : "msg.RS2MS_MsgNotice" 	},
		{ Id : 103, Name : "msg.MS2GW_MsgNotice" 	},
		{ Id : 104, Name : "msg.EntityBase" 	},
		{ Id : 105, Name : "msg.SimpleCounter" 	},
		{ Id : 106, Name : "msg.FreePresentMoney" 	},
		{ Id : 107, Name : "msg.UserWechat" 	},
		{ Id : 108, Name : "msg.UserTask" 	},
		{ Id : 109, Name : "msg.TaskData" 	},
		{ Id : 110, Name : "msg.LuckyDrawItem" 	},
		{ Id : 111, Name : "msg.LuckyDrawRecord" 	},
		{ Id : 112, Name : "msg.ImageData" 	},
		{ Id : 113, Name : "msg.PersonalImage" 	},
		{ Id : 114, Name : "msg.UserBase" 	},
		{ Id : 115, Name : "msg.UserAddress" 	},
		{ Id : 116, Name : "msg.ItemData" 	},
		{ Id : 117, Name : "msg.ItemBin" 	},
		{ Id : 118, Name : "msg.Serialize" 	},
		{ Id : 119, Name : "msg.MS2Server_BroadCast" 	},
		{ Id : 120, Name : "msg.C2GW_AddDeliveryAddress" 	},
		{ Id : 121, Name : "msg.C2GW_DelDeliveryAddress" 	},
		{ Id : 122, Name : "msg.C2GW_ChangeDeliveryAddress" 	},
		{ Id : 123, Name : "msg.GW2C_SendDeliveryAddressList" 	},
		{ Id : 124, Name : "msg.C2GW_ReqRechargeMoney" 	},
		{ Id : 125, Name : "msg.GW2C_RetRechargeMoney" 	},
		{ Id : 126, Name : "msg.C2GW_PlatformRechargeDone" 	},
		{ Id : 127, Name : "msg.GW2C_SendWechatInfo" 	},
		{ Id : 128, Name : "msg.C2GW_StartLuckyDraw" 	},
		{ Id : 129, Name : "msg.GW2C_LuckyDrawHit" 	},
		{ Id : 130, Name : "msg.GW2C_FreePresentNotify" 	},
		{ Id : 131, Name : "msg.GW2C_SendTaskList" 	},
		{ Id : 132, Name : "msg.C2GW_GetTaskReward" 	},
		{ Id : 133, Name : "msg.GW2C_SendLuckyDrawRecord" 	},
		{ Id : 134, Name : "msg.C2GW_ReqTaskList" 	},
		{ Id : 135, Name : "msg.SortInfo" 	},
		{ Id : 136, Name : "msg.C2GW_ReqSort" 	},
		{ Id : 137, Name : "msg.GW2C_RetSort" 	},
		{ Id : 138, Name : "msg.C2GW_ReqGetCash" 	},
		{ Id : 139, Name : "msg.C2GW_RetGetCash" 	},
		{ Id : 140, Name : "msg.C2GW_UpdateUserInfo" 	},
		{ Id : 141, Name : "msg.C2GW_ShareOk" 	},
		{ Id : 142, Name : "msg.CmnRewardInfo" 	},
		{ Id : 143, Name : "msg.GW2C_CmnRewardInfo" 	},
		{ Id : 144, Name : "msg.GW2C_ShareTime" 	}
	]


// Id
export var ProtoIdById : game.Dictionary<table.IProtoIdDefine> = {}
function readProtoIdById(){
  for(let rec of ProtoId) {
    ProtoIdById[rec.Id] = rec; 
  }
}
readProtoIdById();

// Name
export var ProtoIdByName : game.Dictionary<table.IProtoIdDefine> = {}
function readProtoIdByName(){
  for(let rec of ProtoId) {
    ProtoIdByName[rec.Name] = rec; 
  }
}
readProtoIdByName();
}

