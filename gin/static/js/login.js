/*!
 *  Page - Login
 *  @Lastest Update 2016-03-18 12:00
 */
$(document).ready(function(){
    $("#buttonId").click(function(){
        var  nusername = $("#usernameId").val();
        var  password = $("#passwordId").val();
        var  vcode = $("#vcodeId").val();
        var data = "{\"username\":\""+$("#usernameId").val()+"\",\"password\":\""+$("#passwordId").val()
            +"\",\"code\":\""+$("#vcodeId").val()+"\"}";
        $.ajax({
            type: "POST",
            url: "login",
            contentType:"application/json; charset=utf-8",
            dataType: "json",
            data: data,
            success: function(data){
                if (data.code==0){
                    return window.top.location.href="index";
                }else if(data.code==1){
                    alert(JSON.stringify(data));
                    return ;
                }else if (data.code==2){
                    alert(JSON.stringify(data));
                    return ;
                }
            },
            error: function () {
                alert("没有查询到信息！");
            },
        });
    });
});