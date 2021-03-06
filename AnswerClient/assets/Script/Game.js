let Game = {
    async: require('async'),
    moment: require('moment'),
    _: require('lodash'),

    Define: require('./Util/Define'),
    ChickenDefine: require('./Util/ChickenDefine'),
    Tools: require('./Util/Tools'),
    HttpUtil: require('./Util/HttpUtil'),
    PlatformDefine: require('./Util/PlatformDefine'),

    Platform: require('./Platform/Common'),

    AudioController: require('./Controller/AudioController'),
    ConfigController: require('./Controller/ConfigController'),
    GameController: require('./Controller/GameController'),
    LoginController: require('./Controller/LoginController'),
    NetWorkController: require('./Controller/NetWorkController'),
    NotificationController: require('./Controller/NotificationController'),
    ResController: require('./Controller/ResController'),
    TimeController: require('./Controller/TimeController'),

    UserModel: require('./Model/User'),
    RoomModel: require('./Model/Room'),
    TaskModel: require('./Model/Task'),

    GameInstance: null
};

if (cc.sys.platform == cc.sys.WECHAT_GAME) {
    Game.Platform =  require('./Platform/WechatGame');
}

module.exports = Game;

