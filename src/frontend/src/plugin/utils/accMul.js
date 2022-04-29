export function accMul(num, total) {
    return (Math.round(num / total * 10000) / 100.00 + "%");// 小数点后两位百分比
}
export function accMulInt(num, total) {
    return(Math.round(num / total * 10000) / 100.00 / 100.00);
}
// accMul(xx,100)