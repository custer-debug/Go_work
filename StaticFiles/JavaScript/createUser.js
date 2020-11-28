
function validateEmail(email) {
    var re = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}


function SendPostRequest(){
    let check = "Female"
    if(document.getElementById("Male").checked){
        check = "Male";
    }

    const email = document.getElementById("login");
    

    if(!validateEmail(email.value)){
        email.className = "border border-danger rounded"
        return
    }
    email.className = "border border-dark rounded"

    const l = document.getElementById("password").value
    const r = document.getElementById("secondPassword").value
    if(l === r){
        const toSend= {
            firstname : document.getElementById("firstname").value,
            lastname : document.getElementById("lastname").value,
            birthday : document.getElementById("birthday").value,
            gender: check,
            phone : document.getElementById("phone").value,
            login : document.getElementById("login").value,
            password : l
        };
        const json = JSON.stringify(toSend);

        const xhr = new XMLHttpRequest();
        xhr.open("POST","/create");
        xhr.setRequestHeader('Content-Type',"application/json");
        xhr.send(json);
        location.replace("/login");

    }else{
        location.reload();
        alert("Passwords do not match");
    }

}
