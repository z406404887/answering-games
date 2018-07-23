module game {
    export class DataManager {
        static playerModel: PlayerModel;
        static boyDbones: DbonesManage;
        static girlDbones: DbonesManage;

        public static init() {
            DataManager.playerModel = new PlayerModel();
            DataManager.playerModel.RegisterEvent();
            
            DataManager.initDbones();

            table.TBirckInfo = table.TBirckInfo.sort((s1: table.ITBirckInfoDefine, s2: table.ITBirckInfoDefine) => {
                let n1 = splitStringToNumberArray(s1.Info, "-");
                let n2 = splitStringToNumberArray(s2.Info, "-");
                return n1[2] - n2[2];
            })
        }

        public static async initDbones() {
            DataManager.boyDbones = new DbonesManage();
            DataManager.girlDbones = new DbonesManage();

            await DataManager.boyDbones.loadDbones();
            // await DataManager.girlDbones.loadDbones();
        }
    }
}