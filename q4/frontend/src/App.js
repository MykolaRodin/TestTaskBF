import React from 'react';
import logo from './logo.svg';
import './App.css';
import {BrowserRouter as Router, Route, Switch} from 'react-router-dom'
import ListUserComponent from './components/ListUserComponent';
import HeaderComponent from './components/HeaderComponent';
import FooterComponent from './components/FooterComponent';
import CreateUserComponent from './components/CreateUserComponent';

function App() {
  return (
    <div>
        <Router>
              <HeaderComponent />
                <div className="container">
                    <Switch>
                          <Route path = "/" exact component = {ListUserComponent}></Route>
                          <Route path = "/users/:type/:id" component = {ListUserComponent}></Route>
                          <Route path = "/action/:type/:id" component = {CreateUserComponent}></Route>
                         </Switch>
                </div>
              <FooterComponent />
        </Router>
    </div>

  );
}

export default App;
