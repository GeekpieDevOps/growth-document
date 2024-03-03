// fetch: POST /login
// body: { username: 'admin', password: 'admin' }
function login(){
  var username = document.getElementById('username');
  var password = document.getElementById('password');



  fetch('http://10.20.232.78:18080/your-endpoint', {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json'
      },
      body: JSON.stringify({
          username: username.value,
          password: password.value
      })
  })
  .then(response => response.json())
  .then(data => {
      if(data.code == 200){
          // console.log(data.data.type);
          console.log(data.data)
          location.href = 'index.html';
      }
      else{
          console.log(data.message);
      }
  })
  .catch(function(err) {
      console.log(err);
  });
}