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

// 导航区
class Breadcrumb extends React.Component {
    constructor(props) {
        super(props);
        this.state = {local: [{"name": "Home", "id": "abc"}, {"name": "文档", "id": "def"}]};

        // 这边绑定是必要的，这样 `this` 才能在回调函数中使用
        this.handleClick = this.handleClick.bind(this);
    }

    componentDidMount() {
        // console.log("execute right now")
    }

    handleClick(uuid, event) {
        utilsBox.httpRequest("get", "/files/" + uuid).then((resp) => {
            if (resp.status != 200 || resp.data.code != 200) {
                return
            }
            this.setState(prevState => ({
                files: resp.data.data
            }));
        })
    }

    render() {
        const element = this.state.local.map((item) =>
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
        this.state = {files: []};

        // 这边绑定是必要的，这样 `this` 才能在回调函数中使用
        this.handleClick = this.handleClick.bind(this);
    }

    componentDidMount() {
        utilsBox.httpRequest("get", "/files/0").then((resp) => {
            if (resp.status != 200 || resp.data.code != 200) {
                return
            }
            this.setState(prevState => ({
                files: resp.data.data
            }));
        })
    }

    handleClick(id, event) {
        console.log("click" + id)
    }

    handleDoubleClick(uuid, event) {
        utilsBox.httpRequest("get", "/files/" + uuid).then((resp) => {
            if (resp.status != 200 || resp.data.code != 200) {
                return
            }
            this.setState(prevState => ({
                files: resp.data.data
            }));
        })
    }

    render() {
        const element = this.state.files.map((item) =>
            <div className="col"
                 key={item.id}
                 onClick={this.handleClick.bind(this, item.id)}
                 onDoubleClick={this.handleDoubleClick.bind(this, item.uuid)}>
                <div className="card h-100">
                    <img src={"/preview/" + item.uuid} className="card-img-top" alt="..."/>
                    <div className="card-body">
                        <p className="card-text text-truncate">{item.name}</p>
                    </div>
                </div>
            </div>
        );
        return <div className="row row-cols-2 row-cols-sm-3 row-cols-md-4 row-cols-lg-5 row-cols-xl-6 row-cols-xxl-7 g-4">{element}</div>
    }
}

// 渲染
ReactDOM.render(
    <Breadcrumb/>,
    document.getElementById('breadcrumb')
);

ReactDOM.render(
    <FileView/>,
    document.getElementById('fileView')
);
