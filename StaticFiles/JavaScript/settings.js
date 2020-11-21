function logout(){
    location.replace("/logout");
}


function deleteUser(){
location.replace("/delete")
}


if(document.cookie === ""){
    location.replace("/login");
}



function sendRequestProfile(toSend){

    const json = JSON.stringify(toSend)
    const xhr = new XMLHttpRequest();
    xhr.open("POST","/settings");
    xhr.setRequestHeader('Content-Type',"application/json");
    xhr.send(json);
    //alert(json);
}


function createRequestProfile(){
    const Data = {
        firstname:document.getElementById("firstname").value,
        lastname:document.getElementById("lastname").value,
        phone:document.getElementById("phone").value,
        birthday:document.getElementById("birthday").value,
    }
    sendRequestProfile(Data);
}





function createRequestPassword(){
    const Password = {
        oldPassword: document.getElementById("inputPassword").value,
        newPassword1: document.getElementById("inputPassword1").value,
        newPassword2: document.getElementById("inputPassword2").value,
    }
    sendRequestProfile(Password)
}







const button = document.getElementById('button');
const password = document.getElementById('inputPassword');
const password1 = document.getElementById('inputPassword1');
const password2 = document.getElementById('inputPassword2');



button.addEventListener("click", (e) => {
    e.preventDefault();
    checkInputs();
});

function checkInputs(){

    const passwordValue = password.value.trim();
    const password1Value = password.value
    const password2Value = password.value

    var tmp = true;

    if(passwordValue === ""){
        password.parentElement.className = 'form-control is-invalid';
        tmp = false;
    }
    if(password1Value === ""){
        password1.parentElement.className = 'form-control is-invalid';
        tmp = false;

    }
    if(password2Value === ""){
        password2.parentElement.className = 'form-control is-invalid';
        tmp = false;
    }
    if (tmp){ //Create comparison comparison of two strings

        alert("Im here");
    }




}