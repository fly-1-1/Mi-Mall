$(document).ready(function () {
	baseApp.init();
});

let baseApp = {
    init: function () {
        this.initAside()
		this.confirmDelete()
		this.resizeIframe()
        this.changeStatus()
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
    },
    changeStatus:function (){
       $(".chStatus").click(function (){
           let id=$(this).attr("data-id")
           let table=$(this).attr("data-table")
           let field=$(this).attr("data-field")
           let el=$(this)
           $.get("/admin/changeStatus",{id:id,table:table,field:field},function (res){
                console.log(res)
               if (res.success){
                   if (el.attr("src").indexOf("yes")!=-1){
                       el.attr("src","/static/admin/images/no.gif")
                   }else{
                       el.attr("src","/static/admin/images/yes.gif")
                   }
               }
           })
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