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
    checkInputs();
});


function alertError(message){
    
    document.getElementById('var').innerHTML = message;
    document.getElementById("secondPass").className = "alert alert-danger alert-dismissible fade show";

}


function checkInputs(){

    const password = document.getElementById('inputPassword');
    const password1 = document.getElementById('inputPassword1');
    const password2 = document.getElementById('inputPassword2');
    const passwordValue = password.value;
    const password1Value = password1.value;
    const password2Value = password2.value;

    if(passwordValue === ""){
        
        alertError("Enter your password");
        password.className = 'form-control is-invalid';

    }else if(password1Value === ""){
        alertError("Enter new password");
        password1.className = 'form-control is-invalid';
    }else if(password1Value.length < 8){
        console.log("<8");
        alertError("Password must be more than 8 characters ");
      
    }else if(password2Value === ""){

        alertError("Enter second new password");
        password2.className = 'form-control is-invalid';

    }else if(password1Value !== password2Value){
        console.log("password1 !== password2");
        alertError("New passwords is not match");
    }else{
        password.className = password1.className = password2.className = 'form-control is-valid';
        secondPass.className = 'alert alert-danger alert-dismissible fade';

      const Pas = {
          oldPassword:passwordValue,
          newPassword1:password1Value,
          newPassword2:password2Value,
      }
      sendRequest(Pas,Test);
    }

}


function Test(Str){
    console.log(Str)
}




function sendRequest(msgForServer,callback){

    const json = JSON.stringify(msgForServer)
    const xhr = new XMLHttpRequest();
    xhr.open("POST","/settings");

    xhr.onreadystatechange = function (){
    if(xhr.readyState === 4){
        callback(xhr.responseText);
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
    sendRequest(Data,Test);
}