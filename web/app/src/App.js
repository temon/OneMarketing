import React from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';
import LoginPage from './components/auth/LoginPage';
import SignupPage from './components/auth/SignupPage';
// import NotFoundPage from './components/common/NotFoundPage';
import DashboardPage from './components/dashboard/DashboardPage';

function App() {
  return (
      <Router>
        <Switch>
          <Route exact path="/" component={LoginPage} />
          <Route exact path="/signup" component={SignupPage} />
          <Route exact path="/login" component={LoginPage} />
          <Route exact path="/dashboard" component={DashboardPage} />
          {/*<Route component={NotFoundPage} />*/}
        </Switch>
      </Router>
  );
}

export default App;