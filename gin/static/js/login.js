/*!
 *  Page - Login --query
 *  @Lastest Update 2016-03-18 12:00
 */
// $(document).ready(function(){
//     $("#buttonId").click(function(){
//         var  nusername = $("#usernameId").val();
//         var  password = $("#passwordId").val();
//         var  vcode = $("#vcodeId").val();
//         var data = "{\"username\":\""+$("#usernameId").val()+"\",\"password\":\""+$("#passwordId").val()
//             +"\",\"code\":\""+$("#vcodeId").val()+"\"}";
//         $.ajax({
//             type: "POST",
//             url: "login",
//             contentType:"application/json; charset=utf-8",
//             dataType: "json",
//             data: data,
//             success: function(data){
//                 if (data.code==0){
//                     return window.top.location.href="index";
//                 }else if(data.code==1){
//                     alert(JSON.stringify(data));
//                     return ;
//                 }else if (data.code==2){
//                     alert(JSON.stringify(data));
//                     return ;
//                 }
//             },
//             error: function () {
//                 alert("没有查询到信息！");
//             },
//         });
//     });
// });

// var warninghtml={}

var vm = new Vue({
    el: "#app",
    data: {
        username: "",
        password: "",
        warningshow: false
    },
    methods: {
        login: function () {
            var instance = axios.create({
                headers: { 'content-type': 'application/json' }
            });
            // let params = vm.$data
            let params = { username: this.username, password: this.password }
            // console.log("post options: \n", params);
            instance.post("/login", params)
                .then(function (res) {
                    var data = res.data
                    if (data.code == 0) {
                        return window.top.location.href = "index";
                    } else {
                        if (data.code == 1) {
                            vm.warningshow = true
                            $("#warningId").text("用户或密码错误")
                            // vm.showinfo = "用户或密码错误"
                            return;
                        } else if (data.code == 2) {
                            $("#warningId").text("用户或密码错误")
                            // vm.showinfo = "用户或密码错误"
                            return;
                        }
                    }
                })
                .catch(error => {
                    vm.warningshow = true,
                    $("#warningId").text(error)
                    // vm.showinfo = error
                    console.log("axios error:", error);
                });
        }
    }
})