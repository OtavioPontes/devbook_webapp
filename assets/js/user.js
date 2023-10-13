$("#unfollow").on("click", unfollow);

$("#follow").on("click", follow);

$("#edit-profile").on("submit", edit);

$("#edit-password").on("submit", editPassword);

$("#delete-user").on("click", deleteUser);

function unfollow() {
  const userId = $(this).data("user-id");
  $(this).prop("disabled", true);

  $.ajax({
    url: `/users/${userId}/unfollow`,
    method: "POST",
  })
    .done(function () {
      window.location = `/users/${userId}`;
    })
    .fail(function () {
      Swal.fire("Ops...", "Erro ao parar de seguir usuário", "error");
    })
    .always(function () {
      $("#unfollow").prop("disabled", false);
    });
}

function follow() {
  const userId = $(this).data("user-id");
  $(this).prop("disabled", true);

  $.ajax({
    url: `/users/${userId}/follow`,
    method: "POST",
  })
    .done(function () {
      window.location = `/users/${userId}`;
    })
    .fail(function () {
      Swal.fire("Ops...", "Erro ao seguir usuário", "error");
    })
    .always(function () {
      $("#follow").prop("disabled", false);
    });
}

function edit(event) {
  event.preventDefault();

  $(this).prop("disabled", true);

  $.ajax({
    url: "/edit-profile",
    method: "PUT",
    data: {
      name: $("#name").val(),
      email: $("#email").val(),
      nick: $("#nick").val(),
    },
  })
    .done(function () {
      Swal.fire("Sucesso", "Usuário atualizado com sucesso", "success").then(
        function () {
          window.location = "/profile";
        }
      );
    })
    .fail(function () {
      Swal.fire("Ops...", "Erro ao atualizar o perfil", "error");
    })
    .always(function () {
      $("#edit-profile").prop("disabled", false);
    });
}

function editPassword(event) {
  event.preventDefault();

  $(this).prop("disabled", true);

  if ($("#confirm-new-password").val() != $("#new-password").val()) {
    Swal.fire("Ops...", "As novas senhas não estão batendo", "error");
    return;
  }

  $.ajax({
    url: "/edit-password",
    method: "POST",
    data: {
      password: $("#password").val(),
      newPassword: $("#new-password").val(),
    },
  })
    .done(function () {
      Swal.fire("Sucesso", "Senha atualizada com sucesso", "success");
    })
    .fail(function () {
      Swal.fire("Ops...", "Erro ao atualizar a senha", "error");
    })
    .always(function () {
      $("#edit-password").prop("disabled", false);
    });
}

function deleteUser(event) {
  event.preventDefault();

  $(this).prop("disabled", true);

  Swal.fire({
    title: "Atenção!",
    text: "Tem certeza que deseja excluir a sua conta?",
    showCancelButton: true,
    cancelButtonText: "Cancelar",
    icon: "warning",
  }).then(function (confirmation) {
    if (!confirmation.value) return;

    $.ajax({
      url: "/delete-user",
      method: "DELETE",
    })
      .done(function () {
        Swal.fire("Sucesso", "Usuário excluído com sucesso", "success").then(
          function () {
            window.location = "/logout";
          }
        );
      })
      .fail(function () {
        Swal.fire("Ops...", "Erro ao deletar o usuário", "error");
      });
  });
}
