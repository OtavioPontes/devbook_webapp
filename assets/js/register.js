$(`#form-register`).on("submit", createUser);

function createUser(event) {
  event.preventDefault();
  console.log("cadastrando...");

  if ($(`#password`).val() != $(`#confirm-password`).val()) {
    Swal.fire("Ops...", "As senhas devem ser iguais", "error");

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
      Swal.fire(
        "Sucesso!",
        "Usuário Cadastrado com sucesso ✅",
        "success"
      ).then(function () {
        $.ajax({
          url: "/login",
          method: "POST",
          data: {
            email: $("#email").val(),
            password: $("#password").val(),
          },
        })
          .done(function () {
            window.location = "/home";
          })
          .fail(function () {
            Swal.fire("Erro!", "Falha ao autenticar usuário", "error");
          });
      });
    })
    .fail(function () {
      Swal.fire("Erro!", "Erro ao cadastrar usuário 😥", "error");
    });
}
