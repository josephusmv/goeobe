function loginuer(){
    var xhr = new XMLHttpRequest();
    var formData = new FormData();
    formData.append('user', 'Mickey')
    formData.append("pwd", 'Mickey');
    formData.append("expdays", "7");
    xhr.open('POST', "LoginUser", true);
    xhr.onreadystatechange = function () {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            console.log(xhr.responseText);
            showResponse(xhr.responseText);
        }                
    };                   
    xhr.send(formData);
}

function Logout(){
    var xhr = new XMLHttpRequest();
    var formData = new FormData();
    xhr.open('POST', "LogOutUser", true);
    xhr.onreadystatechange = function () {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            console.log(xhr.responseText);
            showResponse(xhr.responseText);
        }                
    };                   
    xhr.send(formData);

}

function showuser(){
    var xhr = new XMLHttpRequest();
    var formData = new FormData();
    xhr.open('POST', "GetCurrentUser", true);
    xhr.onreadystatechange = function () {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            console.log(xhr.responseText);
            showResponse(xhr.responseText);
        }                
    };
    xhr.send(formData);
}

function deleteuser(){
    var xhr = new XMLHttpRequest();
    var formData = new FormData();
    xhr.open('POST', "DeleteUser", true);
    xhr.onreadystatechange = function () {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            console.log(xhr.responseText);
            showResponse(xhr.responseText);
        }                
    };
    xhr.send(formData);
}

function showResponse(responseText) {
    document.getElementById("resultarea").innerText = responseText;
}
