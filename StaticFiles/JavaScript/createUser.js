function SendPostRequest(){
    let check = "Female"
    if(document.getElementById("Male").checked){
        check = "Male";
    }

    console.log(check)

    const l = document.getElementById("password").value
    const r = document.getElementById("secondPassword").value
    if(l === r){
        const toSend= {
            name : document.getElementById("firstname").value,
            surname : document.getElementById("lastname").value,
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
