let utilsBox = function () {

    let httpRequest = function (method, url, data) {
        // TODO :: configs
        axios.defaults.baseURL = 'http://localhost:8000/';
        return axios({
            method: method,
            url: url,
            headers: {'content-type': 'application/x-www-form-urlencoded'},
            data: data
        });
    }

    return {
        httpRequest: function (method, uri, data) {
            return httpRequest(method, uri, data)
        }
    }

}()

// Sidebar
class Sidebar extends React.Component {
    render() {
        return <div className="d-flex flex-column flex-shrink-0 p-3">
            <a href="#" className="d-flex align-items-center mb-3 mb-md-0 me-md-auto link-dark text-decoration-none">
                <svg className="bi me-2" width="40" height="32">
                    <use xlinkHref="#upload"/>
                </svg>
                <span className="fs-4">UPLOAD</span>
            </a>
            <hr/>
            <ul className="nav nav-pills flex-column mb-auto">
                <li className="nav-item">
                    <a href="#" className="nav-link active" aria-current="page">
                        <svg className="bi me-2" width="16" height="16">
                            <use xlinkHref="#hdd"/>
                        </svg>
                        我的硬盘
                    </a>
                </li>
                <li>
                    <a href="#" className="nav-link link-dark">
                        <svg className="bi me-2" width="16" height="16">
                            <use xlinkHref="#hdd"/>
                        </svg>
                        共享硬盘
                    </a>
                </li>
                <hr/>
                <li>
                    <a href="#" className="nav-link link-dark">
                        <svg className="bi me-2" width="16" height="16">
                            <use xlinkHref="#share"/>
                        </svg>
                        与我共享
                    </a>
                </li>
                <li>
                    <a href="#" className="nav-link link-dark">
                        <svg className="bi me-2" width="16" height="16">
                            <use xlinkHref="#clock"/>
                        </svg>
                        最近用过
                    </a>
                </li>
                <li>
                    <a href="#" className="nav-link link-dark">
                        <svg className="bi me-2" width="16" height="16">
                            <use xlinkHref="#star"/>
                        </svg>
                        已加星标
                    </a>
                </li>
                <li>
                    <a href="#" className="nav-link link-dark">
                        <svg className="bi me-2" width="16" height="16">
                            <use xlinkHref="#trash"/>
                        </svg>
                        回收站
                    </a>
                </li>
            </ul>
            <hr/>
        </div>
    }
}

// 导航区
class Breadcrumb extends React.Component {
    constructor(props) {
        super(props);
        this.handleClick = this.handleClick.bind(this);
    }

    handleClick(uuid, event) {
        utilsBox.httpRequest("get", "/files/" + uuid).then((resp) => {
            if (resp.status != 200 || resp.data.code != 200) {
                return
            }
            this.props.onSetState({local: []});
        })
    }

    render() {
        const element = this.props.local.map((item) =>
            <li className="breadcrumb-item" key={item.id}><a className="breadcrumb-link link-dark" href="#" onClick={this.handleClick.bind(this, item.id)}>{item.name}</a></li>
        );
        return (
            <nav aria-label="breadcrumb">
                <ul className="breadcrumb breadcrumb-no-gutter">
                    {element}
                    <li className="breadcrumb-item active" aria-current="page">Billing</li>
                </ul>
            </nav>
        );
    }
}

// 文件区
class FileView extends React.Component {
    constructor(props) {
        super(props);
        this.handleChange = this.handleChange.bind(this);
    }

    componentDidMount() {
        utilsBox.httpRequest("get", "/files/0").then((resp) => {
            if (resp.status != 200 || resp.data.code != 200) {
                return
            }
            this.handleChange(resp.data.data)
        })
    }

    handleChange(files) {
        this.props.onSetState({files: files});
    }

    handleClick(id, event) {
        console.log("click" + id)
    }

    handleDoubleClick(item, event) {
        if (item.mime != "folder") {
            return
        }
        utilsBox.httpRequest("get", "/files/" + item.uuid).then((resp) => {
            if (resp.status != 200 || resp.data.code != 200) {
                return
            }
            this.handleChange(resp.data.data)
        })
    }

    render() {
        const element = this.props.files.map((item) =>
            <div className="col"
                 key={item.id}
                 onClick={this.handleClick.bind(this, item.id)}
                 onDoubleClick={this.handleDoubleClick.bind(this, item)}>
                <div className="card h-100">
                    {item.mime == "folder" ?
                        <img src={"/holder/200x150?text=" + item.name} className="card-img-top" alt="..."/>
                        : <img src={"/preview/" + item.uuid} className="card-img-top" alt="..."/>
                    }
                    <div className="card-body">
                        <p className="card-text text-truncate">{item.name}</p>
                    </div>
                </div>
            </div>
        );
        return <div className="row row-cols-2 row-cols-sm-3 row-cols-md-4 row-cols-lg-5 row-cols-xl-6 row-cols-xxl-7 g-4">{element}</div>
    }
}

// RootView
class RootView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            files: [],
            local: [{"name": "Home", "id": "abc"}, {"name": "文档", "id": "def"}]
        };
        this.handleSetState = this.handleSetState.bind(this);
    }

    handleSetState(json) {
        // this.setState(json);
        this.setState(prevState => (json));
    }

    render() {
        return <div>
            <div className="container-fluid">

                <header className="p-3 mb-3 border-bottom">
                    <div className="container">
                        <div
                            className="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start">
                            <a href="/"
                               className="d-flex align-items-center mb-2 mb-lg-0 text-dark text-decoration-none">
                                <svg className="bi me-2" width="40" height="32" role="img" aria-label="Bootstrap">
                                    <use xlinkHref="#logoSVG"/>
                                </svg>
                            </a>

                            <ul className="nav col-12 col-lg-auto me-lg-auto mb-2 justify-content-center mb-md-0">
                                <li><a href="#" className="nav-link px-2 link-secondary">Overview</a></li>
                                <li><a href="#" className="nav-link px-2 link-dark">Inventory</a></li>
                                <li><a href="#" className="nav-link px-2 link-dark">Customers</a></li>
                                <li><a href="#" className="nav-link px-2 link-dark">Products</a></li>
                            </ul>

                            <form className="col-12 col-lg-auto mb-3 mb-lg-0 me-lg-3">
                                <input type="search" className="form-control" placeholder="Search..."
                                       aria-label="Search"/>
                            </form>

                            <div className="dropdown text-end">
                                <a href="#" className="d-block link-dark text-decoration-none dropdown-toggle"
                                   id="dropdownUser1" data-bs-toggle="dropdown" aria-expanded="false">
                                    <img src="/holder/100x100" alt="mdo" width="32" height="32"
                                         className="rounded-circle"/>
                                </a>
                                <ul className="dropdown-menu text-small" aria-labelledby="dropdownUser1">
                                    <li><a className="dropdown-item" href="#">New project...</a></li>
                                    <li><a className="dropdown-item" href="#">Settings</a></li>
                                    <li><a className="dropdown-item" href="#">Profile</a></li>
                                    <li>
                                        <hr className="dropdown-divider"/>
                                    </li>
                                    <li><a className="dropdown-item" href="#">Sign out</a></li>
                                </ul>
                            </div>
                        </div>
                    </div>
                </header>

                <div className="row">
                    <div className="col-md-2">
                        <Sidebar/>
                    </div>
                    <div className="col-md-10">
                        <Breadcrumb local={this.state.local} onSetState={this.handleSetState}/>
                        <FileView files={this.state.files} onSetState={this.handleSetState}/>
                    </div>
                </div>

                <footer className="d-flex flex-wrap justify-content-between align-items-center py-3 my-4 border-top">
                    <div className="col-md-4 d-flex align-items-center">
                        <a href="/" className="mb-3 me-2 mb-md-0 text-muted text-decoration-none lh-1">
                            <svg className="bi" width="30" height="24">
                                <use xlinkHref="#logoSVG"/>
                            </svg>
                        </a>
                        <span className="text-muted">&copy; 2021 Company, Inc</span>
                    </div>

                    <ul className="nav col-md-4 justify-content-end list-unstyled d-flex">
                        <li className="ms-3"><a className="text-muted" href="#">
                            <svg className="bi" width="24" height="24">
                                <use xlinkHref="#twitter"/>
                            </svg>
                        </a></li>
                        <li className="ms-3"><a className="text-muted" href="#">
                            <svg className="bi" width="24" height="24">
                                <use xlinkHref="#instagram"/>
                            </svg>
                        </a></li>
                        <li className="ms-3"><a className="text-muted" href="#">
                            <svg className="bi" width="24" height="24">
                                <use xlinkHref="#facebook"/>
                            </svg>
                        </a></li>
                    </ul>
                </footer>

            </div>
        </div>
    }
}

// 渲染
ReactDOM.render(
    <RootView/>,
    document.getElementById('root')
);
