import { ApolloClient, ApolloProvider, InMemoryCache } from "@apollo/client";
import { BrowserRouter, Outlet, Route, Routes } from "react-router-dom";
import { Layout } from "./components";
import { AuthProvider } from "./contexts/AuthContext";
import { ClassesPage } from "./pages/classes";
import { ClassPage } from "./pages/classes/:id";
import { InstructorsPage } from "./pages/instructors";
import { InstructorPage } from "./pages/instructors/:id";
import { NewInstructorPage } from "./pages/instructors/new";
import { MembersPage } from "./pages/members";
import { MemberPage } from "./pages/members/:id";
import { MembersClassPage } from "./pages/members_classes/:id";
import { NewMemberPage } from "./pages/members/new";
import { NewClassPage } from "./pages/schedules/:id/classes/new";
import { SignIn } from "./pages/sign-in";
import { SignUp } from "./pages/sign-up";
import { Private, Public } from "./routes";

const client = new ApolloClient({
  uri: `${process.env.REACT_APP_BACKEND_HOST}/query`,
  cache: new InMemoryCache(),
});

export function App() {
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
                  <Layout />
                </Private>
              }
            >
              <Route path="instructors" element={<Outlet />}>
                <Route path="new" element={<NewInstructorPage />} />
                <Route path=":instructorId" element={<InstructorPage />} />
                <Route index element={<InstructorsPage />} />
              </Route>
              <Route path="classes" element={<Outlet />}>
                <Route index element={<ClassesPage />} />
                <Route path=":classId" element={<ClassPage />} />
              </Route>
              <Route
                path="schedules/:id/classes/new"
                element={<NewClassPage />}
              />
              <Route path="members" element={<Outlet />}>
                <Route path="new" element={<NewMemberPage />} />
                <Route index element={<MembersPage />} />
                <Route path=":memberId" element={<MemberPage />} />
              </Route>
              <Route
                path="members_classes/:membersClassId"
                element={<MembersClassPage />}
              />
            </Route>
          </Routes>
        </AuthProvider>
      </BrowserRouter>
    </ApolloProvider>
  );
}
