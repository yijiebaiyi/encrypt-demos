const CryptoJS = require('crypto-js');

const key = "123456789july703";
const iv = CryptoJS.enc.Utf8.parse(key)
function encode(str) {
    str = CryptoJS.enc.Utf8.parse(str);
    return CryptoJS.enc.Base64.stringify(str);
}


function encrypt(word) {
    let encrypted = "";
    const srcs = CryptoJS.enc.Utf8.parse(word)

    encrypted = CryptoJS.AES.encrypt(srcs, iv, {
        iv,
        mode: CryptoJS.mode.CBC,
        paddings: CryptoJS.pad.Pkcs7
    })
    return encrypted.ciphertext.toString()
}



const str = "123456";
const base64str = encode(str);
const res = encrypt(base64str)
console.log("base64: ", base64str)
console.log("aes: ", res)