$(function () {
    loginApp.init();
})
let loginApp = {
    init: function () {
        this.getCaptcha()
        this.imgChange()
    },
    getCaptcha: function () {
        $.get("/admin/captcha?t=" + Math.random(), function (response) {
            console.log(response)
            $("#captchaId").val(response.captchaId)
            $("#captchaImg").attr("src", response.captchaImage)
        })
    },
    imgChange: function () {
        let c = this
        $("#captchaImg").click(function () {
            c.getCaptcha()
        })
    }
}