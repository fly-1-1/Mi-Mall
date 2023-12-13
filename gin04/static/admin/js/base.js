$(document).ready(function () {
	baseApp.init();
});

let baseApp = {
    init: function () {
        this.initAside()
		this.confirmDelete()
		this.resizeIframe()
    },
    initAside: function () {
        $('.aside h4').click(function () {
            $(this).siblings('ul').slideToggle();
        })
    },
	resizeIframe:function(){
		$("#rightMain").height($(window).height()-80)
	},
    confirmDelete: function () {
        $(".delete").click(function (){
			let flag =confirm("您确定要删除吗")
			return flag
		})
    }
}


// $(function(){
//
// 	$('.aside h4').click(function(){
//
// //		$(this).toggleClass('active');
//
// 		$(this).siblings('ul').slideToggle();
// 	})
// })