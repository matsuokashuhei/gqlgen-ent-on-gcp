import { ApolloClient, ApolloProvider, InMemoryCache } from "@apollo/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { AuthProvider } from "./contexts/AuthContext";
import { Home } from "./pages/home";
import { NewInstructorPage } from "./pages/instructors/new";
import { InstructorsPage } from "./pages/instructors";
import { InstructorPage } from "./pages/instructors/:id";
import { SignIn } from "./pages/sign-in";
import { SignUp } from "./pages/sign-up";
import { Public, Private } from "./routes";

const client = new ApolloClient({
  uri: "http://localhost:8080/query",
  cache: new InMemoryCache(),
});

function App() {
  return (
    <ApolloProvider client={client}>
      <BrowserRouter>
        <AuthProvider>
          <Routes>
            <Route
              path="/sign-up"
              element={
                <Public>
                  <SignUp />
                </Public>
              }
            />
            <Route
              path="/sign-in"
              element={
                <Public>
                  <SignIn />
                </Public>
              }
            />
            <Route
              path="/"
              element={
                <Private>
                  <Home />
                </Private>
              }
            />
            <Route
              path="instructors/new"
              element={
                <Private>
                  <NewInstructorPage />
                </Private>
              }
            />
            <Route
              path="instructors/:id"
              element={
                <Private>
                  <InstructorPage />
                </Private>
              }
            />
            <Route
              path="instructors"
              element={
                <Private>
                  <InstructorsPage />
                </Private>
              }
            />
          </Routes>
        </AuthProvider>
      </BrowserRouter>
    </ApolloProvider>
  );
}

export default App;
