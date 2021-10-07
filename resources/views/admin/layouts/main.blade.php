<!DOCTYPE html>
<html lang="en">
<!-- BEGIN HEAD -->
<head>
    <meta charset="utf-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta content="width=device-width, initial-scale=1" name="viewport" />
    <meta name="description" content="Responsive Admin Template" />
    <meta name="author" content="SmartUniversity" />
    <title>Spice Hotel | Bootstrap 4 Admin Dashboard Template + UI Kit</title>
    <!-- google font -->
    <link href="https://fonts.googleapis.com/css?family=Poppins:300,400,500,600,700" rel="stylesheet" type="text/css" />
	<!-- icons -->
    <link href="{{ url('admin') }}/assets/plugins/simple-line-icons/simple-line-icons.min.css" rel="stylesheet" type="text/css" />
    <link href="{{ url('admin') }}/assets/plugins/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css"/>
	<!--bootstrap -->
	<link href="{{ url('admin') }}/assets/plugins/bootstrap/css/bootstrap.min.css" rel="stylesheet" type="text/css" />
	<link href="{{ url('admin') }}/assets/plugins/summernote/summernote.css" rel="stylesheet">
	<!-- morris chart -->
    <link href="{{ url('admin') }}/assets/plugins/morris/morris.css" rel="stylesheet" type="text/css" />
    <!-- Material Design Lite CSS -->
	<link rel="stylesheet" href="{{ url('admin') }}/assets/plugins/material/material.min.css">
	<link rel="stylesheet" href="{{ url('admin') }}/assets/css/material_style.css">
	<!-- animation -->
	<link href="{{ url('admin') }}/assets/css/pages/animate_page.css" rel="stylesheet">
	<!-- Template Styles -->
    <link href="{{ url('admin') }}/assets/css/plugins.min.css" rel="stylesheet" type="text/css" />
    <link href="{{ url('admin') }}/assets/css/style.css" rel="stylesheet" type="text/css" />
    <link href="{{ url('admin') }}/assets/css/responsive.css" rel="stylesheet" type="text/css" />
    <link href="{{ url('admin') }}/assets/css/theme-color.css" rel="stylesheet" type="text/css" />
	<!-- favicon -->
    <link rel="shortcut icon" href="{{ url('admin') }}/assets/img/favicon.ico" /> 
 </head>
 <!-- END HEAD -->
 <body class="page-header-fixed sidemenu-closed-hidelogo page-content-white page-md header-white dark-sidebar-color logo-dark">
    <div class="page-wrapper">
        @include('admin.layouts.header')
        <!-- start page container -->
        <div class="page-container">
            <!-- start sidebar menu -->
            @include('admin.layouts.sidebar')
            <!-- end sidebar menu -->
            <!-- start page content -->
            <div class="page-content-wrapper">
                <div class="page-content">
                    @include('admin.layouts.page-bar')
                    
                    @yield('main')
                </div>
            </div>
            <!-- end page content -->
            <!-- start chat sidebar -->
            @include('admin.layouts.chat-sidebar')
            <!-- end chat sidebar -->
        </div>
        <!-- end page container -->
        @include('admin.layouts.footer')
    </div>
    <!-- start js include path -->
    <script src="{{ url('admin') }}/assets/plugins/jquery/jquery.min.js" ></script>
    <script src="{{ url('admin') }}/assets/plugins/popper/popper.min.js" ></script>
    <script src="{{ url('admin') }}/assets/plugins/jquery-blockui/jquery.blockui.min.js" ></script>
	<script src="{{ url('admin') }}/assets/plugins/jquery-slimscroll/jquery.slimscroll.min.js"></script>
    <!-- bootstrap -->
    <script src="{{ url('admin') }}/assets/plugins/bootstrap/js/bootstrap.min.js" ></script>
    <script src="{{ url('admin') }}/assets/plugins/sparkline/jquery.sparkline.min.js" ></script>
	<script src="{{ url('admin') }}/assets/js/pages/sparkline/sparkline-data.js" ></script>
    <!-- Common js-->
	<script src="{{ url('admin') }}/assets/js/app.js" ></script>
    <script src="{{ url('admin') }}/assets/js/layout.js" ></script>
    <script src="{{ url('admin') }}/assets/js/theme-color.js" ></script>
    <!-- material -->
    <script src="{{ url('admin') }}/assets/plugins/material/material.min.js"></script>
    <!-- animation -->
    <script src="{{ url('admin') }}/assets/js/pages/ui/animations.js" ></script>
    <!-- morris chart -->
    <script src="{{ url('admin') }}/assets/plugins/morris/morris.min.js" ></script>
    <script src="{{ url('admin') }}/assets/plugins/morris/raphael-min.js" ></script>
    <script src="{{ url('admin') }}/assets/js/pages/chart/morris/morris_home_data.js" ></script>
    <!-- end js include path -->
  </body>
</html>