const host = "http://127.0.0.1:4523/m1/4078973-0-default";
const path = {
  signup: host + "/signup",
  login: host + "/login",
  logout: host + "/logout",
  account: host + "/account",
  settings: host + "/settings",
};
const api = {
  signup: host + "/api/v1/user/sign_up",
  login: host + "/api/v1/user/sign_in",
  studentActivities: host + "/api/v1/student/activities",
  signout: (uuid) => host + `/api/v1/user/${uuid}/sign_out`,
  account: (uuid) => host + `/api/v1/user/${uuid}`,
};

export { host, path, api };
