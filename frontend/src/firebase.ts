// import { getAuth, connectAuthEmulator } from "";

// import { initializeApp } from "firebase/app";

export const config = {
  apiKey: process.env.REACT_APP_FIREBASE_API_KEY,
  authDomain: process.env.REACT_APP_FIREBASE_AUTH_DOMAIN,
  projectId: process.env.REACT_APP_FIREBASE_PROJECT_ID,
  storageBucket: process.env.REACT_APP_FIREBASE_STORAGE_BUCKET,
  messagingSenderId: process.env.REACT_APP_FIREBASE_MESSAGING_SENDER_ID,
  appId: process.env.REACT_APP_FIREBASE_APP_ID,
};

// const app = initializeApp(config);

// export const auth = getAuth();
// connectAuthEmulator(auth, "http://localhost:9099");

// import { connectAuthEmulator, getAuth } from "firebase/auth";

// const auth = getAuth();
// if (window.location.hostname === "localhost") {
//   console.log("connectAuthEmulator");
//   connectAuthEmulator(auth, "http://localhost:9099");
// }

// export { auth };
