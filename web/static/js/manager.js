toastr.options.positionClass = 'toast-top-center';

$('#DeviceNameChange').on('show.bs.modal', function (event) {
    const button = $(event.relatedTarget);
    const name = button.data('name');
    let uuid = button.data('uuid');

    const modal = $(this);
    modal.find('.modal-body #name').val(name)
    modal.find('.modal-body #uuid').val(uuid)
});

$('#DeviceError').on('show.bs.modal', function (event) {
    const button = $(event.relatedTarget);
    let uuid = button.data('uuid');

    const url = "/manager/" + uuid + "/error";
    $.get(url,function (data,statue) {
        if (statue === "success") {
            if(data.statue === "success"){
                $("#error").val(data.data)
            }else{
                toastr.error('获取失败')
            }
        }
    });

    const modal = $(this);
    modal.find('.modal-body #error').val(name)
});

$('#changename').click(function () {
    const uuid = $('#uuid').val();
    const newname = $('#name').val();

    const url = "/manager/" + uuid + "/namechange?newname=" + newname;

    $.get(url,function (data,statue) {
        if (statue === "success") {
            if(data.statue === "success"){
                $(location).attr("href","/manager")
            }
        }
        toastr.error('修改失败')
    })
});

$('#DeviceEvent').on('show.bs.modal', function (event) {
    const button = $(event.relatedTarget);
    let uuid = button.data('uuid');

    const url = "/manager/" + uuid + "/event";
    $.get(url,function (data,statue) {
        if (statue === "success") {
            if(data.statue === "success"){
                let signals = JSON.parse(data.data);
                $('#event').val(JSON.stringify(signals, null, "\t"));
                autosize($('#event'))

            }else{
                toastr.error('获取失败')
            }
        }
    });

    const modal = $(this);
    modal.find('.modal-body #error').val(name)
});