toastr.success("如果样式不对，请清理缓存！");
//读取缓存区的数据
var eMV = $("#encryptMobile").val(localStorage.getItem("encryptMobile"));
//初始化复制插件（clipboard）
var clipboard = new ClipboardJS('#btn_copycode');
//按钮事件 提交处理（单次处理）
$("#btn_submit").click(function () {
    var eMV = $("#encryptMobile").val();
    localStorage.setItem("encryptMobile", eMV);//设置缓存区的数据
    var iCV = $("#invitationCode").val().replace(/\r\n/g, "<br/>").replace(/\n/g, "<br/>").replace(/\s/g, " ").replace(" ", "");
    if (eMV == "") { toastr.error("encryptMobile 数据没给我", "对方丑拒你，并嘲讽你"); return; }
    if (iCV == "") { toastr.error("invitationCode 数据没给我", "对方丑拒你，并嘲讽你"); return; }
    iCVS = iCV.split("<br/>");//分割成数组
    iCVS = iCVS.sort();//数组排序
    iCVS = deduplication(iCVS);//本地去重复
    iCVS.forEach(iCVSV => {
        $.post("/ApiOneData", {
            encryptMobile: eMV,
            invitationCode: iCVSV,
        }, function (data) {
            if (data.indexOf("成功") != -1) {
                toastr.success(iCVSV + "结果>" + data, "成功");
                console.log("成功>" + iCVSV + "<结果>" + data);
            } else if (data.indexOf("已经") != -1) {
                toastr.warning(iCVSV + "结果>" + data, "已经帮过");
                console.log("已经帮过>" + iCVSV + "<结果>" + data);
            } else if (data.indexOf("不可以") != -1) {
                toastr.warning(iCVSV + "结果>" + data, "自交么");
                console.log("自交么>" + iCVSV + "<结果>" + data);
            } else {
                toastr.error(iCVSV + "结果>" + data, "失败");
                console.log("失败>" + iCVSV + "<结果>" + data);
            }
        })
    });
    $("#invitationCode").val("");
});
//按钮事件 提交处理（自动获取邀请码并处理，切记请不要关闭页面）
$("#btn_autosubmit").click(function () {
    goAuto();//调用回调函数
    //获取数据
    function goAuto() {
        var eMV = $("#encryptMobile").val();
        localStorage.setItem("encryptMobile", eMV);//设置缓存区的数据
        if (eMV == "") { toastr.error("encryptMobile 数据没给我", "对方丑拒你，并嘲讽你"); return; }
        $("#encryptMobile").attr("readonly", "readonly");//禁用元素
        $("#invitationCode").attr("readonly", "readonly");//禁用元素
        if ($("#invitationCode").val() == "") {
            $.post("/ApiGetCodeData", function (datas) {
                datas = JSON.parse(datas);//转换成Json格式
                var code = "";//存储获取到的数据
                datas.forEach(data => {
                    if (data.code) code += data.code + "\n";
                });
                $("#invitationCode").val(code);
                toastr.warning("三秒后执行", "延时");
                setTimeout(function () {
                    goCode();//调用提交数据
                }, '3000');
            });
        }
    }
    //提交数据
    function goCode() {
        var eMV = $("#encryptMobile").val();
        var iCV = $("#invitationCode").val().replace(/\r\n/g, "<br/>").replace(/\n/g, "<br/>").replace(/\s/g, " ").replace(" ", "");
        if (iCV == "") { toastr.error("invitationCode 数据没给我", "对方丑拒你，并嘲讽你"); return; }
        iCVS = iCV.split("<br/>");//分割成数组
        iCVS = iCVS.sort();//数组排序
        iCVS = deduplication(iCVS);//本地去重复
        iCVS.forEach(iCVSV => {
            $.post("/ApiOneData", {
                encryptMobile: eMV,
                invitationCode: iCVSV,
            }, function (data) {
                if (data.indexOf("成功") != -1) {
                    toastr.success(iCVSV + "结果>" + data, "成功");
                    console.log("成功>" + iCVSV + "<结果>" + data);
                } else if (data.indexOf("已经") != -1) {
                    toastr.warning(iCVSV + "结果>" + data, "已经帮过");
                    console.log("已经帮过>" + iCVSV + "<结果>" + data);
                } else if (data.indexOf("不可以") != -1) {
                    toastr.warning(iCVSV + "结果>" + data, "自交么");
                    console.log("自交么>" + iCVSV + "<结果>" + data);
                } else if (data.indexOf("参与的小伙伴太多") != -1) {
                    toastr.warning("参与的小伙伴太多~", "失败");
                } else {
                    toastr.error(iCVSV + "结果>" + data, "失败");
                    console.log("失败>" + iCVSV + "<结果>" + data);
                }
            })
        });
        $("#invitationCode").val("");
        toastr.warning("五秒后执行", "延时");
        setTimeout(function () {
            goAuto();//重复调用
        }, '5000');
    }
});
//按钮事件 提交邀请码
$("#btn_setcode").click(function () {
    var iCV = $("#set_invitationCode").val().replace(/\r\n/g, "<br/>").replace(/\n/g, "<br/>").replace(/\s/g, " ").replace(" ", "");
    if (iCV == "") { toastr.error("invitationCode 数据没给我", "对方丑拒你，并嘲讽你"); return; }
    iCVS = iCV.split("<br/>");//分割成数组
    iCVS = iCVS.sort();//数组排序
    iCVS = deduplication(iCVS);//本地去重复
    iCVS.forEach(iCVSV => {
        $.post("/ApiSetCodeData", {
            invitationCode: iCVSV,
        }, function (data) {
            if (data.indexOf("成功") != -1) {
                toastr.success(iCVSV + "结果>" + data, "成功");
                console.log("成功>" + iCVSV + "<结果>" + data);
            } else if (data.indexOf("已经") != -1) {
                toastr.warning(iCVSV + "结果>" + data, "已经存在");
                console.log("已经存在>" + iCVSV + "<结果>" + data);
            } else {
                toastr.error(iCVSV + "结果>" + data, "失败");
                console.log("失败>" + iCVSV + "<结果>" + data);
            }
        });
    });
    $("#set_invitationCode").val("");
});
//按钮事件 获取邀请码
$("#btn_getcode").click(function () {
    $.post("/ApiGetCodeData", function (datas) {
        datas = JSON.parse(datas);//转换成Json格式
        var code = "";//存储获取到的数据
        datas.forEach(data => {
            if (data.code) code += data.code + "\n";
        });
        $("#get_invitationCode").val(code);
    });
});
//数据去重复 方法
function deduplication(data) {
    var newArray = new Array();//存储结果的数组
    var tempStr = "";//临时存储的字符串
    data.forEach(datav => {
        if (datav != tempStr) {
            newArray.push(datav);
            tempStr = datav;
        }
    });//本地去重复
    return newArray;
}