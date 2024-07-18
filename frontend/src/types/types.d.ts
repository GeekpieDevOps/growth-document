export declare function notify(
  type: "info" | "alert" | "error",
  msg: string,
  callback: () => void = () => {}
): void;
