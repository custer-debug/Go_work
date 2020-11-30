
function validateEmail(email) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}


function checkLength(doc){
   
    if(doc.value.length < 8){
        doc.className = "form-control is-invalid"
        console.log("Parol is huinya")
        return false
    }
    doc.className = "form-control"
    
return true
}




function SendPostRequest(){

    let check = "Female"
    if(document.getElementById("Male").checked){
        check = "Male";
    }

    const email = document.getElementById("login");
    

    if(!validateEmail(email.value)){
        email.className = "form-control is-invalid"
        return
    }
  
    email.className = "form-control"

    const l = document.getElementById("password")
    const r = document.getElementById("secondPassword")
    
    if(!checkLength(l)){
        return 
    } 
    if(!checkLength(r)){
        return 
    }

    if(l.value === r.value){

        $('#myModal').modal("show") 
        document.getElementById("modalText").innerHTML = "На адрес:<h6> " + email.value + 
        " </h6>отправлено пиcьмо с кодом авторизации.";


        const toSend= {
            firstname : document.getElementById("firstname").value,
            lastname : document.getElementById("lastname").value,
            birthday : document.getElementById("birthday").value,
            gender: check,
            phone : document.getElementById("phone").value,
            login : document.getElementById("login").value,
            password : l.value
        }
        const json = JSON.stringify(toSend);
        const xhr = new XMLHttpRequest();
        xhr.open("POST","/create");
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.send(json);

    }
}


function respCode(response){

    const json = JSON.parse(response);
    console.log(json);


    if(json.Status === "Error"){
        alertModal(json.Subject,json.Body)
        return
    }else if(json.Status === "Success"){
        location.replace("login")

    }


}


function alertModal(subject,body){
    document.getElementById("alertError").className = "alert alert-danger alert-dismissible fade show";
    document.getElementById("subjectError").innerHTML = subject;
    document.getElementById("textError").innerHTML = body;
}



function SendConfirmCode(){

    const code = document.getElementById("code").value;
    if(code.length === 0){
        alertModal("Empty enter","Please enter code")
        return
    }else if(code.length !== 6){
        alertModal("Incorrect code","The code must be 6 characters")
        return
    }



    const msg = "code="+encodeURIComponent(code);
    const xhr = new XMLHttpRequest();
    xhr.open("POST","/create/checkCode");
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');

    xhr.onreadystatechange = function (){
        if(xhr.readyState === 4){
            respCode(xhr.responseText);
            
        }

    }

    xhr.send(msg);


}



