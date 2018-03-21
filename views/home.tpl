<div class="container">
    <h3 class="display-3">Welcome to the Demo Web Application:</h3>
    <p> {{ .message }} </p>
    <div>
    {{ if eq .user "" }}
    To access our valuable content you need to <a href="login">
        <button><span class="glyphicon glyphicon-log-in" aria-hidden="true"></span>&nbsp;Sign in</button>
    </a>
    {{ else }}
        You can proceed to our valuable content <a href="content">
        <button>content</button>
    </a>
    {{ end }}
    </div>
</div>