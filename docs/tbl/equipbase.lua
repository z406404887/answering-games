// Generated by github.com/davyxu/tabtoy
// Version: 2.8.10

module table {
export var Equip : table.IEquipDefine[] = [
		{ Id : 101, Name : "头盔", Desc : "装备获得效果", Pos : 1, SuitId : 0, Sex : 1, LoadPoint : [ "head1_1_00", "head1_1_01", "head1_1_02", "head1_1_03", "head1_1_04" ], Path : "equip_json.101", Skill : [ "1" ], Price : 10, CoinType : 1 	},
		{ Id : 102, Name : "头盔", Desc : "装备获得效果", Pos : 1, SuitId : 0, Sex : 0, LoadPoint : [ "head1_1_00", "head1_1_01", "head1_1_02", "head1_1_03", "head1_1_04" ], Path : "equip_json.103", Skill : [ "1" ], Price : 10, CoinType : 1 	},
		{ Id : 103, Name : "空", Desc : "", Pos : 1, SuitId : 0, Sex : 1, LoadPoint : [  ], Path : "", Skill : [  ], Price : 0, CoinType : 0 	},
		{ Id : 201, Name : "上衣", Desc : "装备获得效果", Pos : 2, SuitId : 0, Sex : 1, LoadPoint : [ "body1_1_00", "body1_1_01", "body1_1_02", "body1_1_03", "body1_1_04" ], Path : "equip_json.201", Skill : [ "2" ], Price : 10, CoinType : 1 	},
		{ Id : 202, Name : "上衣", Desc : "装备获得效果", Pos : 2, SuitId : 0, Sex : 0, LoadPoint : [ "body1_1_00", "body1_1_01", "body1_1_02", "body1_1_03", "body1_1_04" ], Path : "equip_json.203", Skill : [ "2" ], Price : 10, CoinType : 1 	},
		{ Id : 203, Name : "空", Desc : "", Pos : 1, SuitId : 0, Sex : 1, LoadPoint : [  ], Path : "", Skill : [  ], Price : 0, CoinType : 0 	},
		{ Id : 301, Name : "裤子", Desc : "装备获得效果", Pos : 3, SuitId : 0, Sex : 1, LoadPoint : [ "trousers1_1_00", "trousers1_1_01", "trousers1_1_02", "trousers1_1_03" ], Path : "equip_json.301", Skill : [ "3" ], Price : 10, CoinType : 1 	},
		{ Id : 302, Name : "裤子", Desc : "装备获得效果", Pos : 3, SuitId : 0, Sex : 0, LoadPoint : [ "trousers1_1_00", "trousers1_1_01", "trousers1_1_02", "trousers1_1_03" ], Path : "equip_json.303", Skill : [ "3" ], Price : 10, CoinType : 1 	},
		{ Id : 303, Name : "空", Desc : "", Pos : 1, SuitId : 0, Sex : 1, LoadPoint : [  ], Path : "", Skill : [  ], Price : 0, CoinType : 0 	},
		{ Id : 401, Name : "鞋子", Desc : "装备获得效果", Pos : 4, SuitId : 0, Sex : 1, LoadPoint : [ "shoes1_1_00", "shoes1_1_01" ], Path : "equip_json.401", Skill : [ "4" ], Price : 10, CoinType : 1 	},
		{ Id : 402, Name : "鞋子", Desc : "装备获得效果", Pos : 4, SuitId : 0, Sex : 0, LoadPoint : [ "shoes1_1_00", "shoes1_1_01" ], Path : "equip_json.403", Skill : [ "4" ], Price : 10, CoinType : 1 	},
		{ Id : 403, Name : "空", Desc : "", Pos : 1, SuitId : 0, Sex : 1, LoadPoint : [  ], Path : "", Skill : [  ], Price : 0, CoinType : 0 	}
	]


// Id
export var EquipById : game.Dictionary<table.IEquipDefine> = {}
function readEquipById(){
  for(let rec of Equip) {
    EquipById[rec.Id] = rec; 
  }
}
readEquipById();
}

