function LoginUserAjax(user, pwd){   
    var xhr = new XMLHttpRequest();
    var formData = new FormData();
    formData.append('username', user)
    formData.append("password", pwd);
    formData.append("expdays", "7");
    xhr.open('POST', "LoginUser", true);
    xhr.onreadystatechange = function () {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            alert(xhr.responseText);
            DocumentOnLoadNavBar();
        }                
    };                   
    xhr.send(formData);  
}

function LogoutUser(){       
    var xhr = new XMLHttpRequest();
    var formData = new FormData();
    xhr.open('POST', "LogoutUser", true);
    xhr.onreadystatechange = function () {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            //alert(xhr.responseText);
            //DocumentOnLoadNavBar();
            window.location.replace("/");
        }                
    };                   
    xhr.send(formData);  
}                        

//$("#userInput"), $("#pwdInput"), int
function DoLoginUser(userInput, pwdInput, expireDays, handler, errHandler){
    var user = userInput.val();
    var pwd = pwdInput.val();
    
    $.post("LoginUser",
    {
        username: user,
        password: pwd,
        expdays: expireDays
    }, function(data, statues) {
        console.log("Data: " + data + "\nStatus: " + status);
        
        if (statues != "success" || data == null || typeof(data) == undefined) {
            errHandler();
            return;
        }
        
        if (data.singlerow != null) {
            handler(data.singlerow.retLoginUser);
            return;
        }
        
        var obj = jQuery.parseJSON(data);
        if  (statues != "success" 
        || obj.singlerow == null || typeof(obj.singlerow) == undefined) {
            errHandler();
            return;
        }
        var retLoginUser = obj.singlerow.retLoginUser;
        console.log("logged in user: " + retLoginUser);
        handler(retLoginUser);
    });
}


    function GetLoginUser(showNoUserLogin, showLoggedInNames){
        var formData = new FormData();    
        var xhr = new XMLHttpRequest();
        xhr.open('POST', "ShowCurrentUser", true);
        xhr.onreadystatechange = function () {
            if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
                var myObj = JSON.parse(this.responseText);
                var user = myObj.singlerow.retCurrentUserName;

                if (user == "") {
                    showNoUserLogin();
                } else {
                    showLoggedInNames(user);
                }
            }
        };
        xhr.send(formData);
    }