$("#new-post").on("submit", createPost);

$("#edit-post").on("click", updatePost);

$(".delete-post").on("click", deletePost);

$(document).on("click", ".like-post", likePost);
$(document).on("click", ".dislike-post", dislikePost);

function createPost(event) {
  event.preventDefault();

  $.ajax({
    url: "/posts",
    method: "POST",
    data: {
      title: $("#title").val(),
      content: $("#content").val(),
    },
  })
    .done(function () {
      Swal.fire("Sucesso!", "Publicação criada com sucesso", "success").then(
        function () {
          window.location = "/home";
        }
      );
    })
    .fail(function () {
      Swal.fire("Erro!", "Falha na criação da publicação", "error").then(
        function () {
          window.location = "/home";
        }
      );
    });
}

function likePost(event) {
  event.preventDefault();

  const element = $(event.target);

  const postId = element.closest("div").data("post-id");

  element.prop("disabled", true);

  $.ajax({
    url: `/posts/${postId}/like`,
    method: "POST",
  })
    .done(function () {
      const likeCounter = element.next("span");
      const qtdLikes = parseInt(likeCounter.text());
      likeCounter.text(qtdLikes + 1);

      element.addClass("dislike-post");
      element.addClass("text-danger");
      element.removeClass("like-post");
    })
    .fail(function () {
      Swal.fire("Erro!", "Falha ao curtir publicação", "error");
    })
    .always(function () {
      element.prop("disabled", false);
    });
}

function dislikePost(event) {
  event.preventDefault();

  const element = $(event.target);

  const postId = element.closest("div").data("post-id");

  element.prop("disabled", true);

  $.ajax({
    url: `/posts/${postId}/dislike`,
    method: "POST",
  })
    .done(function () {
      const likeCounter = element.next("span");
      const qtdLikes = parseInt(likeCounter.text());
      likeCounter.text(qtdLikes - 1);

      element.removeClass("dislike-post");
      element.removeClass("text-danger");
      element.addClass("like-post");
    })
    .fail(function () {
      Swal.fire("Erro!", "Falha ao descurtir publicação", "error");
    })
    .always(function () {
      element.prop("disabled", false);
    });
}

function updatePost(event) {
  event.preventDefault();
  const postId = $(this).data("post-id");
  $(this).prop("disabled", true);

  $.ajax({
    url: `/posts/${postId}`,
    method: "PUT",
    data: {
      title: $("#title").val(),
      content: $("#content").val(),
    },
  })
    .done(function () {
      Swal.fire(
        "Sucesso!",
        "Publicação atualizada com sucesso",
        "success"
      ).then(function () {
        window.location = "/home";
      });
    })
    .fail(function () {
      Swal.fire("Erro!", "Falha na edição da publicação", "error");
    })
    .always(function () {
      $(this).prop("disabled", false);
    });
}

function deletePost(event) {
  event.preventDefault();

  Swal.fire({
    title: "Atenção!",
    text: "Tem certeza que deseja excluir a publicação?",
    showCancelButton: true,
    cancelButtonText: "Cancelar",
    icon: "warning",
  }).then(function (confirmation) {
    if (!confirmation.value) return;
    const element = $(event.target);

    const post = element.closest("div");
    const postId = post.data("post-id");

    element.prop("disabled", true);

    $.ajax({
      url: `/posts/${postId}`,
      method: "DELETE",
    })
      .done(function () {
        post.fadeOut("slow", function () {
          $(this).remove();
        });
      })
      .fail(function () {
        Swal.fire("Erro!", "Falha ao deletar publicação", "error");
      })
      .always(function () {
        element.prop("disabled", false);
      });
  });
}
