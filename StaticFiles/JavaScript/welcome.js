function logout(){
    location.replace("/logout");

}

if(document.cookie === "") {
    location.replace("/login");
}
