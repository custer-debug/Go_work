const url = "/login"
let inputForm = document.getElementById("formData")

function ParseResponse(response){
    if(response === "Error"){


        //Here message about incorrect login or password

    }else if(response === "Success"){
        location.replace("/welcome")
    }
}


inputForm.addEventListener("submit", (e) => {
    e.preventDefault();
    const formdata = new FormData(inputForm)
    console.log(formdata)

    fetch(url,{
        method:"POST",
        body:formdata,
    }).then(
        async response =>  ParseResponse(await response.text())
    )



})
