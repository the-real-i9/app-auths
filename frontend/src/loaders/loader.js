import { redirect } from "react-router-dom"

export async function homeLoader() {
  const jwt = JSON.parse(localStorage.getItem("i9codesoauth_jwt"))


  if (!jwt) {
    return redirect("/auth")
  }

  return jwt
}

export async function socialAuthLoader({ request }) {
  const url = new URL(request.url);
  const provider = url.searchParams.get("provider");

  switch (provider) {
    case "google":
      // goto http://localhost:5000/auth/google_oauth endpoint in backend and retrieve authorization url
      // redirect to authorization url
      console.log("this is google auth")
      return redirect("/auth") //test
    case "github":
      // goto http://localhost:5000/auth/github_oauth endpoint in backend and retrieve authorization url
      // redirect to authorization url
      console.log("this is github auth")
      return redirect("/auth") // test
    default:
      break;
  }

  return null
}