

//Three functions check cookie
function logout(){
    location.replace("/logout");
}


function deleteUser(){
location.replace("/delete")
}


if(document.cookie === ""){
    location.replace("/login");
}



btnChangePassword = document.getElementById('btnChangePassword');
btnChangePassword.addEventListener("click", (e) => {
    e.preventDefault();
    checkInputsPassword();
});

btnProfileAlert = document.getElementById('btnProfileAlert');
btnProfileAlert.addEventListener("click", (e) => {
    e.preventDefault();
    createRequestProfile();
});




//Function check password changing
function checkInputsPassword(){
    console.log("checkInputsPassword");

    const ID = "idAlertPas";
    const msgAlert = "msgAlert";
    const password = document.getElementById('inputPassword');
    const password1 = document.getElementById('inputPassword1');
    const password2 = document.getElementById('inputPassword2');
    const passwordValue = password.value;
    const password1Value = password1.value;
    const password2Value = password2.value;

    if(passwordValue === ""){
        
        alertError("Enter your password", msgAlert,ID);
        password.className = 'form-control is-invalid';

    }else if(password1Value === ""){

        password.className = 'form-control';
        alertError("Enter new password", msgAlert,ID);
        password1.className = 'form-control is-invalid';
    }else if(password1Value.length < 8){
        console.log("<8");
        alertError("Password must be more than 8 characters ", msgAlert,ID);
      
    }else if(password2Value === ""){
        password1.className = 'form-control';

        alertError("Enter second new password", msgAlert,ID);
        password2.className = 'form-control is-invalid';

    }else if(password1Value !== password2Value){
        console.log("password1 !== password2");
        alertError("New passwords is not match", msgAlert,ID);
    }else{
        password.className = password1.className = password2.className = 'form-control is-valid';
        idAlertPas.className = 'alert alert-danger alert-dismissible fade';

      const Pas = {
          oldPassword:  passwordValue,
          newPassword1: password1Value,
          newPassword2: password2Value,
      }

      sendRequest(Pas,jsonHandler,msgAlert,ID);
    }

}

//Message about error
function alertError(message,idMsg, idAlert){
    
    document.getElementById(idMsg).innerHTML = message;
    document.getElementById(idAlert).className = "alert alert-danger alert-dismissible fade show";

}


//Message about success operation
function alertSuccess(message,idMsg, idAlert){
    
    document.getElementById(idMsg).innerHTML = message;
    document.getElementById(idAlert).className = "alert alert-success alert-dismissible fade show";

}


//Func for create Request profile changing
function createRequestProfile(){
    console.log("createRequestProfile");

    const Data = {
        firstname:document.getElementById("firstname").value,
        lastname:document.getElementById("lastname").value,
        phone:document.getElementById("phone").value,
        birthday:document.getElementById("birthday").value,
    }

    sendRequest(Data,jsonHandler,"msg","ProfileAlert");

}







//Handler func response's server
function jsonHandler(Str, isms,alert){
    console.log(Str);

    if(Str === ""){
        alertError("Server's response is empty",isms,alert);
        return 
    }

    const json = JSON.parse(Str);
    console.log(typeof json.Status);

  
    if(json.Status === "Error"){
        alertError(json.Body,isms,alert);
        return
    }
     
    alertSuccess(json.Body,isms,alert);

   

}



//Function send request on server
function sendRequest(msgForServer,callback,isms, alert){
    console.log(isms,alert);
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





