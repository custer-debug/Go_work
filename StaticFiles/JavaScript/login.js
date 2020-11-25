if(document.cookie !== ""){
        location.replace("/profile");
    }


const url = "/login"
let inputForm = document.getElementById("formData")


function ParseResponse(response){
    if (response === "Error") {
        doc = document.getElementById("message")
        const control = doc.parentElement;
        control.className = "alert alert-danger alert-dismissible fade show";

    } else if (response === "Success") {
        login.className = 'form-control is-valid'
        password.className = 'form-control is-valid'
        location.replace("/profile")
    }
}


function checkInputs(){

    var res = true;

    if(login.value === ""){
        login.className = 'form-control is-invalid';
        res = false;
    }
    if(password.value === ""){
        password.className = 'form-control is-invalid';
        res = false;
    
    }

return res;
}


inputForm.addEventListener("submit", (e) => {
    e.preventDefault();


    if (checkInputs()){
        const formdata = new FormData(inputForm)
        console.log(formdata)
    
        fetch(url,{
            method:"POST",
            body:formdata,
        }).then(
            async response =>  ParseResponse(await response.text())
        )

    }   

})

