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
        this.state = {date: new Date()};

        // 这边绑定是必要的，这样 `this` 才能在回调函数中使用
        this.handleClick = this.handleClick.bind(this);
    }

    componentDidMount() {
        // console.log("execute right now")
    }

    handleClick(name, event) {
        console.log(name)
        this.setState(prevState => ({
            date: new Date()
        }));
        console.log(this.state)
        utilsBox.httpRequest("get", "/files/xxx?id=123654", {"aaa": "bbb"})
    }

    render() {
        return (
            <ul className="nav">
                <li className="nav-item">
                    <a className="nav-link active" href="#" onClick={this.handleClick.bind(this, "Name-A")}>目录A</a>
                </li>
                <li className="nav-item">
                    <a className="nav-link" href="#" onClick={this.handleClick.bind(this, "Name-B")}>目录B</a>
                </li>
                <li className="nav-item">
                    <a className="nav-link disabled">目录C {this.state.date.toLocaleTimeString()}</a>
                </li>
            </ul>
        );
    }
}

// 文件区
class FileView extends React.Component {
    constructor(props) {
        super(props);
        this.state = {date: new Date(), files: []};

        // 这边绑定是必要的，这样 `this` 才能在回调函数中使用
        this.handleClick = this.handleClick.bind(this);
    }

    componentDidMount() {
        utilsBox.httpRequest("get", "/files/xxx?id=0").then((resp) => {
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
            <figure key={item.id} className="figure"
                    onClick={this.handleClick.bind(this, item.id)}
                    onDoubleClick={this.handleDoubleClick.bind(this, item.uuid)}>
                <img src="/holder/200x200" className="figure-img img-fluid rounded"/>
                <figcaption className="figure-caption">{item.name}</figcaption>
            </figure>
        );
        return <div className="col-md-12">{element}</div>
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
