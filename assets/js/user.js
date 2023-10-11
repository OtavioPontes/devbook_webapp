$("#unfollow").on("click", unfollow);

$("#follow").on("click", follow);

$("#edit-profile").on("submit", edit);

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
    url: `/edit-profile`,
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
