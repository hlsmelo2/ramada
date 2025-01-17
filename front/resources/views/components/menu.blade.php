<nav class="navbar navbar-expand-lg bg-body-tertiary">
  <div class="container-fluid">
    <a class="navbar-brand" href="#"><img src="{{ asset('img/logo.png') }}" alt="logo"></a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav">
        @if (!isAuthed())
        <li class="nav-item">
          <a class="nav-link" href="{{ route('login') }}">Login</a>
        </li>
        @endif

        @if (isAuthed())
          <li class="nav-item">
            <a class="nav-link" href="{{ route('logout') }}">Logout</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="{{ route('page.user.create') }}">Cadastrar</a>
          </li>
          <li class="nav-item">
            <a class="nav-link active" aria-current="page" href="{{ route('products') }}">Home</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="{{ route('users') }}">Usuários</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="{{ route('products') }}">Produtos</a>
          </li>
        @endif
      </ul>
    </div>
  </div>
</nav>
