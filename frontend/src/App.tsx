import { ApolloClient, ApolloProvider, InMemoryCache } from "@apollo/client";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import { AuthProvider } from "./contexts/AuthContext";
import { Home } from "./pages/home";
import { NewInstructorPage } from "./pages/instructors/new";
import { InstructorsPage } from "./pages/instructors";
import { ClassesPage } from "./pages/classes";
import { InstructorPage } from "./pages/instructors/:id";
import { SignIn } from "./pages/sign-in";
import { SignUp } from "./pages/sign-up";
import { Public, Private } from "./routes";
import { NewClassPage } from "./pages/schedules/:id/classes/new";
import { ClassPage } from "./pages/classes/:id";

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
            <Route
              path="classes"
              element={
                <Private>
                  <ClassesPage />
                </Private>
              }
            />
            <Route
              path="classes/:id"
              element={
                <Private>
                  <ClassPage />
                </Private>
              }
            />
            <Route
              path="schedules/:id/classes/new"
              element={
                <Private>
                  <NewClassPage />
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
