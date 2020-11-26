

function logout(){
    location.replace("/logout");
}


function deleteUser(){
location.replace("/delete")
}


if(document.cookie === ""){
    location.replace("/login");
}



button.addEventListener("click", (e) => {
    e.preventDefault();
    checkInputsPassword();
});


function alertError(message,idMsg, idAlert){
    
    document.getElementById(idMsg).innerHTML = message;
    document.getElementById(idAlert).className = "alert alert-danger alert-dismissible fade show";

}

function alertSuccess(message,idMsg, idAlert){
    
    document.getElementById(idMsg).innerHTML = message;
    document.getElementById(idAlert).className = "alert alert-success alert-dismissible fade show";

}


function checkInputsPassword(){
    const AD = "alertDanger";
    const ID = "secondPass";
    const password = document.getElementById('inputPassword');
    const password1 = document.getElementById('inputPassword1');
    const password2 = document.getElementById('inputPassword2');
    const passwordValue = password.value;
    const password1Value = password1.value;
    const password2Value = password2.value;

    if(passwordValue === ""){
        
        alertError("Enter your password", AD,ID);
        password.className = 'form-control is-invalid';

    }else if(password1Value === ""){

        password.className = 'form-control';
        alertError("Enter new password", AD,ID);
        password1.className = 'form-control is-invalid';
    }else if(password1Value.length < 8){
        console.log("<8");
        alertError("Password must be more than 8 characters ", AD,ID);
      
    }else if(password2Value === ""){
        password1.className = 'form-control';

        alertError("Enter second new password", AD,ID);
        password2.className = 'form-control is-invalid';

    }else if(password1Value !== password2Value){
        console.log("password1 !== password2");
        alertError("New passwords is not match", AD,ID);
    }else{
        password.className = password1.className = password2.className = 'form-control is-valid';
        secondPass.className = 'alert alert-danger alert-dismissible fade';

      const Pas = {
          oldPassword:  passwordValue,
          newPassword1: password1Value,
          newPassword2: password2Value,
      }
      sendRequest(Pas,CallbackFunc,AD,ID);
    }

}


function CallbackFunc(Str, isms,alert){
    console.log(Str);
    const json = JSON.parse(Str);
    console.log(typeof json.Status);

   if(json === ""){
    alertError("Server's response is empty",isms,alert);
   }else{

    if(json.Status === "Error"){
        alertError(json.Body,isms,alert);
    }
     
    alertSuccess(json.Body,isms,alert);

   }

}




function sendRequest(msgForServer,callback,isms, alert){

    const json = JSON.stringify(msgForServer)
    const xhr = new XMLHttpRequest();
    xhr.open("POST","/settings");

    xhr.onreadystatechange = function (){
    if(xhr.readyState === 4){

        callback(xhr.responseText,isms,alert);
        }
    }

    xhr.setRequestHeader('Content-Type',"application/json");
    xhr.send(json);
}





function createRequestProfile(){
    const Data = {
        firstname:document.getElementById("firstname").value,
        lastname:document.getElementById("lastname").value,
        phone:document.getElementById("phone").value,
        birthday:document.getElementById("birthday").value,
    }
    sendRequest(Data,CallbackFunc);
}