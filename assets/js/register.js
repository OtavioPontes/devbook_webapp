$(`#form-register`).on("submit", createUser);

function createUser(event) {
  event.preventDefault();
  console.log("cadastrando...");

  if ($(`#password`).val() != $(`#confirm-password`).val()) {
    alert("Senhas devem ser iguais");
    return;
  }

  $.ajax({
    url: "/users",
    method: "POST",
    data: {
      name: $(`#name`).val(),
      email: $(`#email`).val(),
      nick: $(`#nick`).val(),
      password: $(`#password`).val(),
    },
  })
    .done(function () {
      alert("Usuário Cadastrado com sucesso ✅");
    })
    .fail(function () {
      alert("Erro ao cadastrar usuário 😥");
    });
}
