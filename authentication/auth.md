This project's goal is to implement various auth flows, like basic, jwt and oauth. the storing mechanism in the ui is always the same: save a httpOnly cookie with an authorization jwt. This balances out security (CSRF, XSS), simplicity of implementation, and performance . It only covers the fundamentals: if you expect some deep dives into a complete auth flow with account recovery, email notification service, multiple emails and so on, please visit (https://github.com/E-nkv/Auth-G).


