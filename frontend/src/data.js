const host = "http://localhost:3000";
const path = {
  signup: "/signup",
  login: "/login",
  logout: "/logout",
  account: "/account",
  settings: "/settings",
};
const api = {
  signup: "/api/v1/user/sign_up",
  login: "/api/v1/user/sign_in",
  signout: (uuid) => `/api/v1/user/${uuid}/sign_out`,
  account: (uuid) => `/api/v1/user/${uuid}`,
};

export { host, path, api };
