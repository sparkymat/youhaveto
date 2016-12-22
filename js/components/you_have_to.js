import React from 'react';

export default class YouHaveTo extends React.Component {
  constructor(props) {
    super(props);

    this.state = {};
  }

  render() {
    return (
      <div>
        <nav className="navbar navbar-dark bg-primary">
          <div className="row">
            <a className="navbar-brand col-md-2" href="#" style={ {marginRight: 0} }><h3>YouHaveTo</h3></a>
            <div className="col-md-10">
              <input type="text" id="search" placeholder="Search for movie, book or song to recommend..." className="form-control form-control-lg" autoFocus={ true } />
            </div>
          </div>
        </nav>
      </div>
    )
  }
}
