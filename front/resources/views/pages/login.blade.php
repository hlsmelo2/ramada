@extends('.template.default')

@section('body')
    <h1>Login</h1>
    @include('components.alert', ['flashKey' => 'loginFailed', 'type' => 'danger'])
    @include('components.alert', ['flashKey' => 'userCreate'])

    <form action="{{ route('try.login') }}" method="POST">
        @csrf()
        <div class="input-group mb-3">
            <span class="input-group-text" id="basic-addon1">Email</span>
            <input type="text" name="email" class="form-control" value="user1@example.com" placeholder="Email" aria-label="Email" aria-describedby="basic-addon1">
        </div>

        <div class="input-group mb-3">
            <span class="input-group-text" id="basic-addon1">Password</span>
            <input type="password" name="password" class="form-control" value="Password@123" placeholder="Password" aria-label="Password" aria-describedby="basic-addon1">
        </div>

        <a href="{{ route('page.user.create') }}" class="d-block">Cadastrar</a>
        <button type="submit" class="btn btn-primary">Login</button>
    </form>
@endsection
