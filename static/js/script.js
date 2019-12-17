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