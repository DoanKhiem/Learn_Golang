@extends('admin.layouts.main')

@section('main')

            
            <!-- start widget -->
            <div class="state-overview">
                <div class="row">
                    <div class="col-xl-3 col-md-6 col-12">
                        <div class="info-box bg-blue">
                        <span class="info-box-icon push-bottom"><i class="material-icons">style</i></span>
                        <div class="info-box-content">
                            <span class="info-box-text">Phòng ban</span>
                            <span class="info-box-number">450</span>
                            <div class="progress">
                            <div class="progress-bar width-60"></div>
                            </div>
                            <span class="progress-description">
                                60% Increase in 28 Days
                                </span>
                        </div>
                        <!-- /.info-box-content -->
                        </div>
                        <!-- /.info-box -->
                    </div>
                    <!-- /.col -->
                    <div class="col-xl-3 col-md-6 col-12">
                        <div class="info-box bg-orange">
                        <span class="info-box-icon push-bottom"><i class="material-icons">card_travel</i></span>
                        <div class="info-box-content">
                            <span class="info-box-text">Chức vụ</span>
                            <span class="info-box-number">155</span>
                            <div class="progress">
                            <div class="progress-bar width-40"></div>
                            </div>
                            <span class="progress-description">
                                40% Increase in 28 Days
                                </span>
                        </div>
                        <!-- /.info-box-content -->
                        </div>
                        <!-- /.info-box -->
                    </div>
                    <!-- /.col -->
                    <div class="col-xl-3 col-md-6 col-12">
                        <div class="info-box bg-purple">
                        <span class="info-box-icon push-bottom"><i class="material-icons">phone_in_talk</i></span>
                        <div class="info-box-content">
                            <span class="info-box-text">Hợp đồng</span>
                            <span class="info-box-number">52</span>
                            <div class="progress">
                            <div class="progress-bar width-80"></div>
                            </div>
                            <span class="progress-description">
                                80% Increase in 28 Days
                                </span>
                        </div>
                        <!-- /.info-box-content -->
                        </div>
                        <!-- /.info-box -->
                    </div>
                    <!-- /.col -->
                    <div class="col-xl-3 col-md-6 col-12">
                        <div class="info-box bg-success">
                        <span class="info-box-icon push-bottom"><i class="material-icons">monetization_on</i></span>
                        <div class="info-box-content">
                            <span class="info-box-text">Nhân viên</span>
                            <span class="info-box-number">13,921</span><span>$</span>
                            <div class="progress">
                            <div class="progress-bar width-60"></div>
                            </div>
                            <span class="progress-description">
                                60% Increase in 28 Days
                                </span>
                        </div>
                        <!-- /.info-box-content -->
                        </div>
                        <!-- /.info-box -->
                    </div>
                    <!-- /.col -->
                    </div>
                </div>
            <!-- end widget -->
            <div class="row">
                <div class="col-md-12">
                    <div class="card card-topline-green">
                        <div class="card-head">
                            <header>Markers</header>
                            <div class="tools">
                                <a class="fa fa-repeat btn-color box-refresh" href="javascript:;"></a>
                                <a class="t-collapse btn-color fa fa-chevron-down" href="javascript:;"></a>
                                <a class="t-close btn-color fa fa-times" href="javascript:;"></a>
                            </div>
                        </div>
                        <div class="card-body ">
                            <div id="gmap_marker" class="gmaps"> </div>
                        </div>
                    </div>
                </div>
            </div>

@endsection

        


