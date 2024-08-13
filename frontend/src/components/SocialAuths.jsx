import { useSubmit } from "react-router-dom"

function SocialAuths() {
  const submit = useSubmit()

  return (
    <>
      <button
        className="google-login-btn"
        onClick={() => {
          submit({ provider: "google" }, { method: "GET", action: "/auth" })
        }}
      >
        Login with Google
      </button>
      <button
        className="github-login-btn"
        onClick={() => {
          submit({ provider: "github" }, { method: "GET", action: "/auth" })
        }}
      >
        Login with GitHub
      </button>
    </>
  )
}

export default SocialAuths
